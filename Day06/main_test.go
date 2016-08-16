/* Author: Mecu Sorin       Phone: 0040747020102 */

package Day06

import (
	"testing"

	"github.com/gonum/matrix/mat64"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

func TestBootstaping(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Santa's fire hazard Suite")
}

var _ = Describe("Follow Santa's instructions for lightning a house", func() {
	DescribeTable("Creating an action should return an error when invalid coordinates are provided",
		func(r1, c1, r2, c2 int) {
			_, err := newAction(toggle, r1, c1, r2, c2)
			Expect(err).ShouldNot(Succeed())
		},
		Entry("Invalid row start", -1, 1, 1, 1),
		Entry("Invalid col start", 1, -1, 1, 1),
		Entry("Invalid row end", 1, 1, 1+maxRows, 1),
		Entry("Invalid col end", 1, 1, 1, 1+maxCols),
		Entry("Invalid row1,row2 order", 21, 1, 1, 1),
		Entry("Invalid col1,col2 order", 1, 21, 1, 1),
	)

	Specify("Should understand Santa's command format", func() {
		_, err := parseAction("turn on 80,957 through 776,968")
		Expect(err).Should(Succeed())
		_, err = parseAction("toggle 277,130 through 513,244")
		Expect(err).Should(Succeed())
		_, err = parseAction("turn off 62,266 through 854,434")
		Expect(err).Should(Succeed())
		_, err = parseAction("turn offf 62,266 through 854,434")
		Expect(err).ShouldNot(Succeed())

		_, err = parseAction("turn offf ,266 through 854,434")
		Expect(err).ShouldNot(Succeed())
		_, err = parseAction("turn off 22,266 through 854,4.34")
		Expect(err).ShouldNot(Succeed())

	})

	Context("Given Puzzle 1 verbs", func() {
		BeforeEach(func() {
			switchToDialect1()
		})
		Specify("Actions should work", func() {
			testGrid := ledGrid{mat64.NewDense(maxRows, maxCols, nil)}
			// 00
			// 00
			Expect(mat64.Sum(testGrid)).To(Equal(0.0))

			// turnOn
			action, err := newAction(turnOn, 0, 0, 499, 999)
			Expect(err).Should(Succeed())
			testGrid.apply(action)
			expectedSum := float64(action.rows() * action.cols())
			// 10
			// 10
			Expect(mat64.Sum(testGrid)).To(Equal(expectedSum))

			// toggle
			action, err = newAction(toggle, 0, 0, 999, 499)
			Expect(err).Should(Succeed())
			testGrid.apply(action)
			// 01
			// 10
			Expect(mat64.Sum(testGrid)).To(Equal(expectedSum))

			// turnOff
			action, err = newAction(turnOff, 500, 0, 999, 499)
			Expect(err).Should(Succeed())
			testGrid.apply(action)
			// 00
			// 10
			Expect(mat64.Sum(testGrid)).To(Equal(expectedSum / 2.))
		})
	})

	Context("Given Puzzle 2 verbs", func() {
		BeforeEach(func() {
			switchToDialect2()
		})
		Specify("Actions should work", func() {
			testGrid := ledGrid{mat64.NewDense(maxRows, maxCols, nil)}
			// 00
			// 00
			Expect(mat64.Sum(testGrid)).To(Equal(0.0))

			// turnOn
			action, err := newAction(turnOn, 0, 0, 499, 999)
			Expect(err).Should(Succeed())
			testGrid.apply(action)
			expectedSum := float64(action.rows() * action.cols())
			// 10
			// 10
			Expect(mat64.Sum(testGrid)).To(Equal(expectedSum))

			// turnOff
			action, err = newAction(turnOff, 0, 0, 999, 499)
			Expect(err).Should(Succeed())
			testGrid.apply(action)
			// 00
			// 10
			Expect(mat64.Sum(testGrid)).To(Equal(expectedSum / 2.))

			// toggle
			action, err = newAction(toggle, 0, 500, 999, 999)
			Expect(err).Should(Succeed())
			testGrid.apply(action)
			// 00
			// 32
			Expect(mat64.Sum(testGrid)).To(Equal(5. * expectedSum / 2.))
		})
	})
})
