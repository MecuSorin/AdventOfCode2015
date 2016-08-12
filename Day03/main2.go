/* Author: Mecu Sorin       Phone: 0040747020102 */

package main

import (
	"fmt"
	"unicode/utf8"
)

func process(instructions string, santas ...*santa) error {
	fifoSantaChannel := make(chan *santa, len(santas))
	for _, santa := range santas {
		fifoSantaChannel <- santa
	}
	defer close(fifoSantaChannel)
	radio := make(chan rune)
	go func() {
		defer close(radio)
		for i, width := 0, 0; i < len(instructions); i += width {
			command, runeWidth := utf8.DecodeRuneInString(instructions[i:])
			width = runeWidth
			radio <- command
		}
	}()

	for santa := range fifoSantaChannel {
		command, ok := <-radio
		if !ok {
			break
		}
		if err := santa.makeMove(command); nil != err {
			return err
		}
		fifoSantaChannel <- santa
	}
	return nil
}

func getHousesVisited(santas ...*santa) (int, error) {
	if len(santas) < 1 {
		return 0, fmt.Errorf("No Santas were provided to get visited houses count")
	}
	allHouses := make(map[position]int, len(santas[0].history))
	for i := 0; i < len(santas); i++ {
		for p, v := range santas[i].history {
			allHouses[p] += v
		}
	}
	return len(allHouses), nil
}
