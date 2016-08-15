/* Author: Mecu Sorin       Phone: 0040747020102 */

package day02

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"

	"testing"
)

func TestBootstaping(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Elves Suite")
}

var _ = Describe("Elves should know:", func() {
	Describe("How to interpret the gift sizes", func() {
		It("Should return an error on an improper gift format like 1,2,3", func() {
			_, _, err := getWrappingMaterials("1,2,3")
			Expect(err).ShouldNot(Succeed())
		})
		It("Should interpret a gift format like 1x2x3", func() {
			gifts, err := readGifts("1x2x3")
			Expect(err).Should(Succeed())
			Expect(len(gifts)).To(Equal(1))
		})
		It("Should interpret a gift format like 1x2x3\n5x6x7\n", func() {
			gifts, err := readGifts("1x2x3\n5x6x7\n")
			Expect(err).Should(Succeed())
			Expect(len(gifts)).To(Equal(2))
		})
	})

	DescribeTable("How much paper they need for wrapping the gifts",
		func(w, h, l, expectedSurface int) {
			g := gift{w: w, h: h, l: l}
			Expect(g.getWrappingPaper()).To(Equal(expectedSurface))
		},
		Entry("Gift 1", 2, 3, 4, 58),
		Entry("Gift 2", 1, 1, 10, 43),
	)

	DescribeTable("How much ribbon they need for wrapping the gifts",
		func(w, h, l, expectedRibbon int) {
			g := gift{w: w, h: h, l: l}
			Expect(g.getRibbonNeededForWrapping()).To(Equal(expectedRibbon))
		},
		Entry("Gift ribbon 1", 2, 3, 4, 34),
		Entry("Gift ribbon 2", 1, 1, 10, 14),
	)
})
