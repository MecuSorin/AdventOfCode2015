/* Author: Mecu Sorin       Phone: 0040747020102 */

package Day08

import (
	"bufio"
	"os"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/onsi/ginkgo/extensions/table"
)

func TestBootstaping(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Santa's digital list - counting chars Suite")
}

var _ = Describe("Counting chars", func() {

	DescribeTable("Countig samples",
		func(text string, expectedDelta int) {
			charsCount := charsCountDelta(text)
			Expect(charsCount).To(Equal(expectedDelta))
		},

		Entry("empty string ", `""`, 2),
		Entry(`"abc"`, `"abc"`, 2),
		Entry(`"aaa\"aaa"`, `"aaa\"aaa"`, 3),
		Entry(`"\x27"`, `"\x27"`, 5),
		Entry(`"\\"`, `"\\"`, 3),
		Entry(`"yy\"\"uoao\"uripabop"`, `"yy\"\"uoao\"uripabop"`, 5),
		Entry(`"byc\x9dyxuafof\\\xa6uf\\axfozomj\\olh\x6a"`, `"byc\x9dyxuafof\\\xa6uf\\axfozomj\\olh\x6a"`, 14),
		Entry(`"zf\x23\\hlj\\kkce\\d\\asy\"yyfestwcdxyfj"`, `"zf\x23\\hlj\\kkce\\d\\asy\"yyfestwcdxyfj"`, 10),
		Entry(`"\\\\mouqqcsgmz"`, `"\\\\mouqqcsgmz"`, 4),
		Entry(`"\\\x33\\\""`, `"\\\x33\\\""`, 8),
	)

	DescribeTable("Encoding samples",
		func(text, expectedEncoded string) {
			Expect(encodeString(text)).To(Equal(expectedEncoded))
		},
		Entry(`a`, `a`, `"a"`),
		Entry(`"`, `"`, `"\""`),
		Entry(`\`, `\`, `"\\"`),
		Entry(`""`, `""`, `"\"\""`),
		Entry(`"abc"`, `"abc"`, `"\"abc\""`),
		Entry(`"aaa\"aaa"`, `"aaa\"aaa"`, `"\"aaa\\\"aaa\""`),
		Entry(`"\x27"`, `"\x27"`, `"\"\\x27\""`),
	)

	Context("Puzzle tests", func() {
		DescribeTable("Given the input",
			func(alteration func(string) string, expectation int) {
				puzzleInput, err := os.Open("puzzle.input")
				Expect(err).Should(Succeed())
				defer puzzleInput.Close()

				counter := 0
				scanner := bufio.NewScanner(puzzleInput)
				for scanner.Scan() {
					text := alteration(scanner.Text())
					counter += charsCountDelta(text)
				}
				Expect(scanner.Err()).Should(Succeed())

				Expect(counter).To(Equal(expectation))
			},

			Entry("Puzzle 1", func(x string) string { return x }, 1350),
			Entry("Puzzle 2", func(x string) string { return encodeString(x) }, 2085),
		)
	})
})
