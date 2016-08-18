/* Author: Mecu Sorin       Phone: 0040747020102 */

package Day07

import (
	"strings"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var (
	givenSample = `123 -> x
456 -> y
x AND y -> d
x OR y -> e
x LSHIFT 2 -> f
y RSHIFT 2 -> g
NOT x -> h
NOT y -> i`
	_ = Describe("Parser should work", func() {
		DescribeTable("Individual operations should work",
			func(instructions string, assertions func(state map[string]int)) {
				state, err := process(strings.Split(instructions, "\n"))
				Expect(err).Should(Succeed())
				assertions(state)
			},

			Entry("SET", "123 -> x", func(state map[string]int) {
				Expect(state["x"]).To(Equal(123))
			}),
			Entry("NOT", "NOT 123 -> x", func(state map[string]int) {
				Expect(state["x"]).To(Equal(65412))
			}),
			Entry("AND", "5 AND 3 -> x", func(state map[string]int) {
				Expect(state["x"]).To(Equal(1))
			}),
			Entry("OR", "5 OR 3 -> x", func(state map[string]int) {
				Expect(state["x"]).To(Equal(7))
			}),
			Entry("LSHIFT", "5 LSHIFT 2 -> x", func(state map[string]int) {
				Expect(state["x"]).To(Equal(20))
			}),
			Entry("RSHIFT", "20RSHIFT 2 -> x", func(state map[string]int) {
				Expect(state["x"]).To(Equal(5))
			}),
		)

		Specify("Sample", func() {
			state, err := process(strings.Split(givenSample, "\n"))
			Expect(err).Should(Succeed())
			Expect(state["d"]).To(Equal(72))
			Expect(state["e"]).To(Equal(507))
			Expect(state["f"]).To(Equal(492))
			Expect(state["g"]).To(Equal(114))
			Expect(state["h"]).To(Equal(65412))
			Expect(state["i"]).To(Equal(65079))
			Expect(state["x"]).To(Equal(123))
			Expect(state["y"]).To(Equal(456))
		})

		Specify("Should stop at cycles", func() {
			_, err := process(strings.Split("NOT x->y", "\n"))
			Expect(err).ShouldNot(Succeed())
		})

		Specify("getValue shoud accept only lexonNumber or lexonVariable", func() {
			_, err := getValue(lexon{lexonKeywordAND, 3, "AND"}, make(map[string]int))
			Expect(err).ShouldNot(Succeed())
		})

		Specify("identifyParser should return error on invalid/unknown grammar", func() {
			lexer := newLexer("NOT x->y")
			_, err := identifyParser([]parser{parserSET, parserAND, parserOR, parserLSHIFT, parserRSHIFT}, lexer.collect())
			Expect(err).ShouldNot(Succeed())
		})

		Specify("identifyParser should return error on bad lexonNumber", func() {
			lexer := lexer{emitter: make(chan lexon, 4)}
			lexer.emitter <- lexon{lexonKeywordNOT, 0, "NOT"}
			lexer.emitter <- lexon{lexonNumber, 4, "0x1AZ"}
			lexer.emitter <- lexon{lexonVariable, 6, "ab3"}
			lexer.emitter <- lexon{lexonKeywordLSHIFT, 10, "LSHIFT"}
			close(lexer.emitter)
			_, err := identifyParser(knownParsers, lexer.collect())
			Expect(err).ShouldNot(Succeed())
		})

		Specify("Should skip empty lines", func() {
			state, err := process(strings.Split("NOT 1->x\n\nNOT x->y", "\n"))
			Expect(err).Should(Succeed())
			Expect(state["y"]).To(Equal(1))
		})

		Specify("Should skip empty lines", func() {
			_, err := process(strings.Split("NOT 1->x\n 1 NOT x->y", "\n"))
			Expect(err).ShouldNot(Succeed())
		})
	})
)
