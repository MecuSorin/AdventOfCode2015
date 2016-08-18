/* Author: Mecu Sorin       Phone: 0040747020102 */

package Day07

import (
	"strings"
	"unicode"
)

// Pos represents the position in string
type Pos int

type lexer struct {
	emitter chan lexon
	input   string
	start   Pos
	pos     Pos
	width   Pos
}

type stateFn func(l *lexer) stateFn

func newLexer(input string) *lexer {
	result := &lexer{emitter: make(chan lexon), input: input}
	go result.run()
	return result
}

// run is the lexer state machine engine
func (l *lexer) run() {
	for state := lexAction; nil != state; {
		state = state(l)
	}
	close(l.emitter)
}

func lexError(l *lexer) stateFn {
	l.emit(lexonError)
	//fmt.Println(l.input[l.start:])
	return nil
}

func lexGateEndTerminal(l *lexer) stateFn {
	l.pos = l.start + Pos(len(gateEndTerminal))
	l.emit(lexonKeywordEndGateTerminal)
	return lexAction
}

func lexNumber(l *lexer) stateFn {
	l.acceptRun("0123456789")
	l.emit(lexonNumber)
	return lexAction
}

func lexIdentifier(l *lexer) stateFn {
loop:
	for {
		switch r := l.next(); {
		case isAlphaNumeric(r):
		default:
			break loop
		}
	}
	l.backup()
	word := l.input[l.start:l.pos]
	wordLexonType, ok := grammarKeywords[word]
	if ok {
		l.emit(wordLexonType)
		return lexAction
	}
	l.emit(lexonVariable)
	return lexAction
}

func lexAction(l *lexer) stateFn {
	for {
		switch r := l.next(); {
		case r == eof || isEndOfLine(r):
			return nil
		case unicode.IsDigit(r):
			l.backup()
			return lexNumber
		case isSpace(r):
			l.ignore()
		case unicode.IsLetter(r):
			l.backup()
			return lexIdentifier
		case r == '-':
			if strings.HasPrefix(l.input[l.start:], gateEndTerminal) {
				l.backup()
				return lexGateEndTerminal
			}
			return lexError
		default:
			return lexError
		}
	}
}
