/* Author: Mecu Sorin       Phone: 0040747020102 */

package Day05

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

func TestBootstaping(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Santa is searching for nice/naughty texts Suite")
}

var _ = Describe("Santa should know how to categorize words", func() {
	DescribeTable("When a word is provided",
		func(word string, expectedNiceWord bool) {
			actual := categorize(word)
			Expect(actual).To(Equal(expectedNiceWord))
		},
		Entry("Should recognize if nice sample 1", "ugknbfddgicrmopn", true),
		Entry("Should recognize if nice sample 2", "aaa", true),
		Entry("Should recognize if naughty sample 1", "jchzalrnumimnmhp", false),
		Entry("Should recognize if naughty sample 2", "haegwjzuvuyypxyu", false),
		Entry("Should recognize if naughty sample 3", "dvszwmarrgswjxmb", false),
	)
	Specify("Should find words containing 2 consecutive letters", func() {
		Expect(doubleLetter.err).Should(Succeed())

		p := createMatcher(doubleLetter)
		Expect(p("aa")).To(Equal(true))
		Expect(p("baa")).To(Equal(true))
		Expect(p("baa")).To(Equal(true))
		Expect(p("aab")).To(Equal(true))
		Expect(p("xx")).To(Equal(true))
		Expect(p("abcdde")).To(Equal(true))
		Expect(p("aabbccdd")).To(Equal(true))
		Expect(p("")).To(Equal(false))
		Expect(p("abab")).To(Equal(false))

	})

	Specify("Should find words containing at least 3 vowels", func() {
		Expect(specificVocals.err).Should(Succeed())

		p := createMatcher(specificVocals)
		Expect(p("aei")).To(Equal(true))
		Expect(p("xazegov")).To(Equal(true))
		Expect(p("aeiouaeiouaeiou")).To(Equal(true))
		Expect(p("aertrtrtrt")).To(Equal(false))
	})

	Specify("Should identify naughty words", func() {
		Expect(naughtyStrings.err).Should(Succeed())

		p := createMatcher(naughtyStrings)
		Expect(p("abei")).To(Equal(false))
		Expect(p("xacdov")).To(Equal(false))
		Expect(p("aeiouaeiouaeipq")).To(Equal(false))
		Expect(p("xy")).To(Equal(false))

		Expect(p("xay")).To(Equal(true))
	})
})
