/* Author: Mecu Sorin       Phone: 0040747020102 */

package main

import "unicode/utf8"

func countFloorsAndGiveTheFirstCharThatBringsSantaOnDesiredFloor(instructions string, desiredFloor int) (int, <-chan int) {
	floor := 0
	runeCounter := 0
	firstCharToTheFloor := make(chan int, 1)
	defer close(firstCharToTheFloor)

	if floor == desiredFloor {
		select {
		case firstCharToTheFloor <- runeCounter:
		default:
		}
	}

	for i, width := 0, 0; i < len(instructions); i += width {
		rune, runeWidth := utf8.DecodeRuneInString(instructions[i:])
		width = runeWidth
		switch rune {
		case '(':
			floor++
		case ')':
			floor--
		}
		runeCounter++
		if floor == desiredFloor {
			select {
			case firstCharToTheFloor <- runeCounter:
			default:
			}
		}
	}
	return floor, firstCharToTheFloor
}
