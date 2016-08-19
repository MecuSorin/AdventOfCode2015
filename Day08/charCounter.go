/* Author: Mecu Sorin       Phone: 0040747020102 */

package Day08

import "regexp"

var (
	counterPattern      = regexp.MustCompile(`(?P<Ascii>\\x[0-9A-Fa-f]{2})|(?P<Backslash>\\{2})|(?P<Quote>\\\")`)
	encodeStringPattern = regexp.MustCompile(`(?P<Backslash>\\)|(?P<Quote>\")`)
)

func countMatches(text string) int {
	rez := counterPattern.FindAllStringSubmatch(text[:len(text)-1], -1)
	if nil == rez {
		return 0
	}
	over := 0
	for i := range rez {
		for j, name := range counterPattern.SubexpNames() {
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

func encodeString(text string) string {
	return "\"" + encodeQuotesAndBackslashes(text) + "\""
}

func encodeQuotesAndBackslashes(text string) string {
	rez := encodeStringPattern.FindAllStringSubmatchIndex(text, -1)
	if nil == rez {
		return text
	}
	start := 0
	encoded := ""
	for i := range rez {
		encoded = encoded + text[start:rez[i][0]]
		start = rez[i][1]
		groupsNames := encodeStringPattern.SubexpNames()
		for j := range groupsNames {
			if 0 == j || -1 == rez[i][j*2] {
				continue
			}
			switch groupsNames[j] {
			case "Backslash":
				encoded = encoded + "\\\\"
			case "Quote":
				encoded = encoded + "\\\""
			}
			break
		}
	}
	encoded = encoded + text[start:]
	return encoded
}
