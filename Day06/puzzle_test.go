/* Author: Mecu Sorin       Phone: 0040747020102 */

package Day06

import (
	"bufio"
	"os"

	"github.com/gonum/matrix/mat64"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Follow Santa's instructions for lightning a house", func() {
	Specify("First puzzle", func() {
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
		Expect(mat64.Sum(houseLEDGrid)).To(Equal(543903.))
	})

})
