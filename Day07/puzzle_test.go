/* Author: Mecu Sorin       Phone: 0040747020102 */

package Day07

import (
	"bufio"
	"os"

	. "github.com/onsi/ginkgo"
	//. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var _ = Describe("Santa needs help assembling the circuit", func() {
	Context("Bobby's kit's instructions booklet", func() {
		BeforeEach(func() {
			getInitialState = defaultGetInitialState
			hookSETUpdating = defaultHookSETUpdating
		})
		AfterEach(func() {
			getInitialState = defaultGetInitialState
			hookSETUpdating = defaultHookSETUpdating
		})

		processPuzzleInput := func(assertOnState func(map[string]int)) {
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
			assertOnState(state)
		}
		Specify("Puzzle 1", func() {
			processPuzzleInput(func(state map[string]int) {
				aVal, ok := state["a"]
				Expect(ok).To(BeTrue())
				Expect(aVal).To(Equal(956))
			})
		})
		Specify("Puzzle 2", func() {
			hookSETUpdating = func(s map[string]int, l lexon, v int) {
				if "b" == l.val {
					s["b"] = 956
					return
				}
				defaultHookSETUpdating(s, l, v)
			}
			processPuzzleInput(func(state map[string]int) {
				aVal, ok := state["a"]
				Expect(ok).To(BeTrue())
				Expect(aVal).To(Equal(40149))
				/*
					//anotherState := make(map[string]int)
					//anotherState["b"] = state["a"]
					//getInitialState = func() map[string]int { return anotherState }

					 take the signal you got on wire a, override wire b to that signal,
					   and reset the other wires (including wire a). What new signal is ultimately provided to wire a

					newB := state["a"]
					for k := range state {
						state[k] = 0
					}
					state["b"] = newB
					getInitialState = func() map[string]int { return state }
					processPuzzleInput(func(newState map[string]int) {
						aVal, ok := newState["a"]
						Expect(ok).To(BeTrue())
						Expect(aVal).To(Equal(0))
					})
				*/
			})

		})

	})
})
