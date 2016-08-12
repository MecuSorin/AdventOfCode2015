/* Author: Mecu Sorin       Phone: 0040747020102 */

package day03

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"

	"testing"
)

func TestBootstaping(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Santa's travels Suite")
}

var _ = Describe("Santa travel on drunk elf instructions", func() {
	Context("Single Santa:", func() {
		DescribeTable("Given move directions checks position",
			func(housedVisited int, moveCommands string, x, y int) {
				s := newSanta()
				Expect(s.move(moveCommands)).Should(Succeed())
				Expect(len(s.history)).To(Equal(housedVisited))
				Expect(s.pos).To(Equal(position{x: x, y: y}))
			},

			Entry("On ^ goes up", 2, "^", 0, 1),
			Entry("On v goes down", 2, "v", 0, -1),
			Entry("On < goes left", 2, "<", -1, 0),
			Entry("On > goes right", 2, ">", 1, 0),
			Entry("On ^>v< goes back to start", 4, "^>v<", 0, 0),
			Entry("On ^v<> goes up", 3, "^v<>", 0, 0),
		)

		DescribeTable("Given move directions checks houses visited",
			func(moveCommands string, housesVisited int) {
				s := newSanta()
				Expect(s.move(moveCommands)).Should(Succeed())
				Expect(len(s.history)).To(Equal(housesVisited))
			},

			Entry("On > to have 2 houses visited", "^", 2),
			Entry("On ^>v< to have 4 houses visited", "^>v<", 4),
			Entry("On ^v^v^v^v^v to have 2 houses visited", "^v^v^v^v^v", 2),
		)
	})

	Context("More Santas:", func() {
		Specify("Error should be returned if no Santa is provided", func() {
			_, err := getHousesVisited()
			Expect(err).ShouldNot(Succeed())
		})

		Specify("Counting visited houses from 1 Santas should work", func() {
			santa1 := newSanta()
			houses, _ := getHousesVisited(santa1)
			Expect(houses).To(Equal(1))
			santa1.moveDown()
			houses, _ = getHousesVisited(santa1)
			Expect(houses).To(Equal(2))
			santa1.moveUp()
			houses, _ = getHousesVisited(santa1)
			Expect(houses).To(Equal(2))
		})

		Specify("Counting visited houses from 2 santas should work", func() {
			santa1 := newSanta()
			santa2 := newSanta()
			houses, _ := getHousesVisited(santa1, santa2)
			Expect(houses).To(Equal(1))
			santa1.moveDown()
			houses, _ = getHousesVisited(santa1, santa2)
			Expect(houses).To(Equal(2))
			santa2.moveDown()
			houses, _ = getHousesVisited(santa1, santa2)
			Expect(houses).To(Equal(2))
		})

		DescribeTable("2 Santas should go in different directions",
			func(radioBroadcast string, housesVisited int) {
				Santa := newSanta()
				RoboSanta := newSanta()
				Expect(process(radioBroadcast, Santa, RoboSanta)).Should(Succeed())
				actual, _ := getHousesVisited(Santa, RoboSanta)
				Expect(actual).To(Equal(housesVisited))
			},
			Entry("On no command shoud be 1 house visited", "", 1),
			Entry("On ^> shoud be 3 houses visited", "^>", 3),
			Entry("On ^>v< shoud be 3 houses visited", "^>v<", 3),
			Entry("On ^v^v^v^v^v shoud be 11 houses visited", "^v^v^v^v^v", 11),
		)
	})
})
