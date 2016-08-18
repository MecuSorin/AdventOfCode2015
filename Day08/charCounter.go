/* Author: Mecu Sorin       Phone: 0040747020102 */

package Day08

import "regexp"

var (
	uberPattern = regexp.MustCompile(`(?P<Ascii>\\x[0-9A-Fa-f]{2})|(?P<Backslash>\\{2})|(?P<Quote>\\\")`)
)

func countMatches(text string) int {
	rez := uberPattern.FindAllStringSubmatch(text[:len(text)-1], -1)
	if nil == rez {
		return 0
	}
	over := 0
	for i := range rez {
		for j, name := range uberPattern.SubexpNames() {
			if 0 == j || "" == rez[i][j] {
				continue
			}
			switch name {
			case "Ascii":
				over += 3
			case "Backslash":
				over++
			case "Quote":
				over++
			}
			break
		}

	}
	return over
}

func charsCountDelta(text string) int {
	return 2 + countMatches(text)
}
