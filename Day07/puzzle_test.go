/* Author: Mecu Sorin       Phone: 0040747020102 */

package Day07

import (
	"bufio"
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var _ = Describe("Santa needs help assembling the circuit", func() {
	DescribeTable("Bobby's kit's instructions booklet",
		func(expectedValueForVarA int) {
			puzzleInput, err := os.Open("puzzle.input")
			Expect(err).Should(Succeed())
			defer puzzleInput.Close()

			var instructions []string
			scanner := bufio.NewScanner(puzzleInput)
			for scanner.Scan() {
				instructions = append(instructions, scanner.Text())
			}
			state, err := process(instructions)
			Expect(err).Should(Succeed())
			aVal, ok := state["a"]
			Expect(ok).To(BeTrue())
			Expect(aVal).To(Equal(expectedValueForVarA))

		},

		Entry("Puzzle 1", 956),
		//Entry("using puzzle 2 dialect", switchToDialect2,14687245.),
	)
})
