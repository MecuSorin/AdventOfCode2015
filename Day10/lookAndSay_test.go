/* Author: Mihai-Sorin MECU   0040.747.02.01.02 */

package Day10

import (
	"fmt"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

func TestBootstaping(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Elfs are playing Look And Say Suite")
}

func iterateLookAndSay(seed string, times uint) string {
	for i := times; i > 0; i-- {
		fmt.Printf("%d\t", i)
		seed = lookAndSay(seed)
	}
	return seed
}

var _ = Describe("Given the Look And Say rules", func() {
	DescribeTable("Should comply to the samples",
		func(givenInput, expectedOutput string) {
			Expect(lookAndSay(givenInput)).To(Equal(expectedOutput))
		},
		Entry("Sample 1", "1", "11"),
		Entry("Sample 11", "11", "21"),
		Entry("Sample 21", "21", "1211"),
		Entry("Sample 1211", "1211", "111221"),
		Entry("Sample 111221", "111221", "312211"),
	)
	XContext("Given puzzles", func() {
		var puzzleSeed string
		BeforeEach(func() {
			puzzleSeed = "1113122113"
		})

		Specify("Given the input 1113122113 and iterate LookAndSay 40 times should work properly", func() {
			Expect(len(iterateLookAndSay(puzzleSeed, 40))).To(Equal(360154))
		})
		Specify("Given the input 1113122113 and iterate LookAndSay 50 times should work properly", func() {
			Expect(len(iterateLookAndSay(puzzleSeed, 50))).To(Equal(5103798))
		})

	})
})
