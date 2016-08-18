/* Author: Mecu Sorin       Phone: 0040747020102 */

package Day07

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

func TestBootstaping(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Santa's improper gift Suite")
}

var _ = Describe("The lexer should work", func() {
	var sut *lexer
	Context("Lexer helpers", func() {
		Describe("lexer.next", func() {
			BeforeEach(func() {
				sut = newLexer("a")
			})
			Specify("should return eof if pos is at the end of the input", func() {
				sut.pos = 1
				Expect(sut.next()).To(Equal(eof))
				sut.pos = 2
				Expect(sut.next()).To(Equal(eof))
			})
			Specify("should return a rune if pos is not at the end of the input", func() {
				Expect(sut.next()).To(Equal('a'))
			})
		})

	})
	Context("Lexer engine", func() {
		DescribeTable("Should know how to handle",
			func(input string, expectedTokens []lexon) {
				sut = newLexer(input)
				Expect(sut.collect()).To(Equal(expectedTokens))
			},
			Entry(" empty string", "", []lexon{}),
			Entry(" white spaces", " ", []lexon{}),
			Entry(" 1 ", " 1 ", []lexon{{lexonNumber, 1, "1"}}),
			Entry(" a ", " a ", []lexon{{lexonVariable, 1, "a"}}),
			Entry(" a1 ", " a1 ", []lexon{{lexonVariable, 1, "a1"}}),
			Entry(" AND ", " AND ", []lexon{{lexonKeywordAND, 1, "AND"}}),
			Entry(" -> ", " -> ", []lexon{{lexonKeywordEndGateTerminal, 1, "->"}}),
			Entry(" LSHIFT ", " LSHIFT ", []lexon{{lexonKeywordLSHIFT, 1, "LSHIFT"}}),
			Entry(" NOT ", " NOT ", []lexon{{lexonKeywordNOT, 1, "NOT"}}),
			Entry(" OR ", " OR ", []lexon{{lexonKeywordOR, 1, "OR"}}),
			Entry(" RSHIFT ", " RSHIFT ", []lexon{{lexonKeywordRSHIFT, 1, "RSHIFT"}}),

			Entry("123 ->v", "123 ->v", []lexon{
				{lexonNumber, 0, "123"},
				{lexonKeywordEndGateTerminal, 4, "->"},
				{lexonVariable, 6, "v"},
			}),
			Entry("NOT12ab3LSHIFT", "NOT12ab1LSHIFT", []lexon{
				{lexonVariable, 0, "NOT12ab1LSHIFT"},
			}),
			Entry("NOT 12ab3 LSHIFT", "NOT 12ab3 LSHIFT", []lexon{
				{lexonKeywordNOT, 0, "NOT"},
				{lexonNumber, 4, "12"},
				{lexonVariable, 6, "ab3"},
				{lexonKeywordLSHIFT, 10, "LSHIFT"},
			}),
			Entry("Fail on unsuported symbols", "a: 2", []lexon{
				{lexonVariable, 0, "a"},
				{lexonError, 1, ":"},
			}),
			Entry("Fail on invalid grammar", " 2 --> a", []lexon{
				{lexonNumber, 1, "2"},
				{lexonError, 3, "-"},
			}),
		)
	})

})
