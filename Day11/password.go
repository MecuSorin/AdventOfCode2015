/* Author: Mecu Sorin       Phone: 0040747020102 */

package Day11

import (
	"strings"

	"unicode/utf8"
)

type santaPassword struct {
	base26
}

func (p santaPassword) String() string {
	return padLeft(p.base26.String(), "a", 8)
}

func getNextPassword(currentPass string) (string, error) {
	pass, err := getBase26FromString(currentPass)
	if nil != err {
		return "", err
	}
	for {
		pass++
		candidate := santaPassword{pass}
		if candidate.has3ConsecutiveLetters() && (!candidate.hasConfusingLetters()) && candidate.hasTwoDifferentLetterPairs() {
			return candidate.String(), nil
		}
	}
}

func (p santaPassword) has3ConsecutiveLetters() bool {
	pass := []byte(p.String())
	for i := len(pass) - 3; i >= 0; i-- {
		if pass[i]+1 == pass[i+1] && pass[i+1]+1 == pass[i+2] {
			return true
		}
	}
	return false
}

func (p santaPassword) hasConfusingLetters() bool {
	return strings.ContainsAny(p.String(), "iol")
}

func (p santaPassword) hasTwoDifferentLetterPairs() bool {
	doubleLetters := map[string]struct{}{}
	pass := p.String()
	for length, i, w, previousRune := len(pass), 0, 0, utf8.RuneError; i < length; i += w {
		r, rw := utf8.DecodeRuneInString(pass[i:])
		w = rw

		if previousRune == r {
			pair := string([]rune{previousRune, r})
			doubleLetters[pair] = struct{}{}
			previousRune = utf8.RuneError
		}
		previousRune = r
	}
	return len(doubleLetters) > 1
}

func padLeft(str, pad string, lenght int) string {
	for {
		str = pad + str
		if len(str) > lenght {
			return str[len(str)-lenght : len(str)]
		}
	}
}
