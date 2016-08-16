/* Author: Mecu Sorin       Phone: 0040747020102 */

package Day06

import (
	"bufio"
	"os"

	"github.com/gonum/matrix/mat64"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var _ = Describe("Follow Santa's instructions for lightning a house", func() {
	DescribeTable("Translating the Santa instructions",
		func(swichToDialect func(), expectedLight float64) {
			swichToDialect()
			samples, err := os.Open("puzzle1.input")
			Expect(err).Should(Succeed())
			defer samples.Close()
			scanner := bufio.NewScanner(samples)
			houseLEDGrid := ledGrid{mat64.NewDense(maxRows, maxCols, nil)}
			for scanner.Scan() {
				action, err := parseAction(scanner.Text())
				Expect(err).Should(Succeed())
				houseLEDGrid.apply(action)
			}
			Expect(scanner.Err()).Should(Succeed())
			Expect(mat64.Sum(houseLEDGrid)).To(Equal(expectedLight))
		},

		Entry("using puzzle 1 dialect", switchToDialect1, 543903.),
		Entry("using puzzle 2 dialect", switchToDialect2,14687245.),
	)
})
