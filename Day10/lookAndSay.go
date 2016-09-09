/* Author: Mihai-Sorin MECU   0040.747.02.01.02 */

package Day10

import (
	"fmt"
	"unicode/utf8"
)

type repeatableRune struct {
	rune  rune
	count uint
}

func (rr repeatableRune) String() string {
	if utf8.RuneError == rr.rune {
		return ""
	}
	return fmt.Sprintf("%d%c", rr.count, rr.rune)
}

func (rr *repeatableRune) processAnotherRune(r rune) (result *repeatableRune, sameRune bool) {
	if rr.rune == r {
		rr.count++
		return rr, true
	}
	return &repeatableRune{r, 1}, false
}

func lookAndSay(text string) string {
	currentLetter := &repeatableRune{utf8.RuneError, 0}
	result := ""
	for i, runeWidth := 0, 0; i < len(text); i += runeWidth {
		r, rWidth := utf8.DecodeRuneInString(text[i:])
		runeWidth = rWidth
		sameRune := false
		oldLetter := currentLetter
		currentLetter, sameRune = currentLetter.processAnotherRune(r)
		if !sameRune {
			result += oldLetter.String()
		}
	}
	result += currentLetter.String()
	return result
}
