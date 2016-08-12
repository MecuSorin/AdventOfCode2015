/* Author: Mecu Sorin       Phone: 0040747020102 */

package day03

import (
	"fmt"
	"unicode/utf8"
)

type position struct {
	x, y int
}

type santa struct {
	pos     position
	history map[position]int
}

func newSanta() *santa {
	result := &santa{history: make(map[position]int)}
	result.addCurrentPositionToHistory()
	return result
}

func (s *santa) makeMove(radioCommand rune) error {
	switch radioCommand {
	case '^':
		s.moveUp()
	case 'v':
		s.moveDown()
	case '<':
		s.moveLeft()
	case '>':
		s.moveRight()
	default:
		return fmt.Errorf("Invalid movement command %q", radioCommand)
	}
	return nil
}

func (s *santa) move(instructions string) error {
	for i, runeWidth := 0, 0; i < len(instructions); i += runeWidth {
		rune, width := utf8.DecodeRuneInString(instructions[i:])
		runeWidth = width
		if err := s.makeMove(rune); nil != err {
			return err
		}
	}
	return nil
}

func (s *santa) moveUp() {
	s.pos = position{x: s.pos.x, y: s.pos.y + 1}
	s.addCurrentPositionToHistory()
}
func (s *santa) moveDown() {
	s.pos = position{x: s.pos.x, y: s.pos.y - 1}
	s.addCurrentPositionToHistory()
}
func (s *santa) moveLeft() {
	s.pos = position{x: s.pos.x - 1, y: s.pos.y}
	s.addCurrentPositionToHistory()
}
func (s *santa) moveRight() {
	s.pos = position{x: s.pos.x + 1, y: s.pos.y}
	s.addCurrentPositionToHistory()
}

func (s *santa) addCurrentPositionToHistory() {
	s.history[s.pos]++
}
