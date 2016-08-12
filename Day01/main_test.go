/* Author: Mecu Sorin       Phone: 0040747020102 */

package main

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

func TestBootstaping(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Santa floor Suite")
}

const neverTouchTheBasement = 99999

var _ = Describe("Santa should know:", func() {
	DescribeTable("first instruction that will get him on the basement",
		func(instructions string, desiredFloor, expected int) {
			_, actualChan := countFloorsAndGiveTheFirstCharThatBringsSantaOnDesiredFloor(instructions, desiredFloor)
			var actual int
			select {
			case position, ok := <-actualChan:
				if ok {
					actual = position
				} else {
					actual = neverTouchTheBasement
				}
			default:
				actual = neverTouchTheBasement
			}
			Expect(actual).To(Equal(expected))
		},
		Entry("Sample 01", ")", -1, 1),
		Entry("Sample 02", "()()))", -1, 5),
		Entry("Sample 03", "()", -1, neverTouchTheBasement),
		Entry("Sample 04", "()", 0, 0),
	)

	DescribeTable("how to follow floor travel instructions",
		func(instructions string, expected int) {
			actual, _ := countFloorsAndGiveTheFirstCharThatBringsSantaOnDesiredFloor(instructions, 0)
			Expect(actual).To(Equal(expected))
		},
		Entry("Sample 01", "()()", 0),
		Entry("Sample 02", "(())", 0),
		Entry("Sample 03", "(((", 3),
		Entry("Sample 04", "(()(()(", 3),
		Entry("Sample 05", "))(((((", 3),
		Entry("Sample 06", "())", -1),
		Entry("Sample 07", "))(", -1),
		Entry("Sample 08", ")))", -3),
		Entry("Sample 09", ")())())", -3),
		Entry("Sample 10", "", 0),
		Entry("The question", "", 280),
	)
})
