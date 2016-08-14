/* Author: Mecu Sorin       Phone: 0040747020102 */

package Day05

import (
	"fmt"
	"regexp"

	"github.com/glenn-brown/golang-pkg-pcre/src/pkg/pcre"
)

//http://stackoverflow.com/questions/35736368/regex-to-match-repeated-characters
var (
	specificVocals = newPattern("(?i)[aeiou].*[aeiou].*[aeiou]", true)
	doubleLetter   = newPatternPCRE("([a-zA-Z])\\1+", true)
	naughtyStrings = newPatternPCRE("(?=ab|cd|pq|xy)", false)
	categorize     = createMatcher(specificVocals, doubleLetter, naughtyStrings)

	doubleDoubleLetterWithoudOverlap    = newPatternPCRE("([a-zA-Z]{2}).*\\1", true)
	repeatableLetterWithSingleLetterGap = newPatternPCRE("([a-zA-Z]).\\1", true)
	categorizePuzzle2                   = createMatcher(doubleDoubleLetterWithoudOverlap, repeatableLetterWithSingleLetterGap)
)

type matcher interface {
	MatchString(string) bool
}
type pcreRegex struct {
	pcre.Regexp
}

func (regex pcreRegex) MatchString(text string) bool {
	matcher := regex.MatcherString(text, 0)
	return matcher.Matches()
}

type pattern struct {
	regex    matcher
	matching bool
	err      error
}

func newPattern(regexPattern string, matching bool) pattern {
	r, e := regexp.Compile(regexPattern)
	return pattern{regex: r, matching: matching, err: e}
}

func newPatternPCRE(regexPattern string, matching bool) pattern {
	r, e := pcre.Compile(regexPattern, 1)
	var err error
	if nil != e {
		err = fmt.Errorf("Fail to compile PCRE pattern:%v", e)
	}
	return pattern{regex: pcreRegex{r}, matching: matching, err: err}
}

func createMatcher(patterns ...pattern) func(string) bool {
	return func(text string) bool {
		for i := range patterns {

			p := patterns[i]
			if p.regex.MatchString(text) != p.matching {
				return false
			}
		}
		return true
	}
}
