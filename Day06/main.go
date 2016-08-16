/* Author: Mecu Sorin       Phone: 0040747020102 */

package Day06

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/gonum/matrix/mat64"
)

type actionVerb int

const (
	_ actionVerb = iota
	turnOn
	turnOff
	toggle
	maxRows = 1000
	maxCols = 1000
)

var (
	regex         = regexp.MustCompile(`(\w+\s*\w+)\s+(\d+),(\d+)\s+through\s+(\d+),(\d+)`)
	puzzleToggle  = puzzle1Toggle
	puzzleTurnOff = puzzle1TurnOff
	puzzleTurnOn  = puzzle1TurnOn
)

type action struct {
	Verb           actionVerb
	r1, r2, c1, c2 int
}

func parseAction(input string) (action, error) {
	matchedGroups := regex.FindStringSubmatch(input)
	if nil == matchedGroups {
		return action{}, fmt.Errorf("Unknown action format: %q", input)
	}
	var verb actionVerb
	switch matchedGroups[1] {
	case "toggle":
		verb = toggle
	case "turn on":
		verb = turnOn
	case "turn off":
		verb = turnOff
	default:
		return action{}, fmt.Errorf("Unknown action format: %q", input)
	}
	// all parsing below is guarded by the regex interpreter because of the pattern used
	r1, _ := strconv.Atoi(matchedGroups[2])
	c1, _ := strconv.Atoi(matchedGroups[3])
	r2, _ := strconv.Atoi(matchedGroups[4])
	c2, _ := strconv.Atoi(matchedGroups[5])
	return newAction(verb, r1, c1, r2, c2)
}

func newAction(verb actionVerb, r1, c1, r2, c2 int) (action, error) {
	if 0 > r1 || maxRows < r2 || r1 > r2 || 0 > c1 || maxCols < c2 || c1 > c2 {
		return action{}, fmt.Errorf("Invalid coordinates provided")
	}
	return action{Verb: verb, r1: r1, r2: r2, c1: c1, c2: c2}, nil
}

func (a action) rows() int {
	return a.r2 - a.r1 + 1
}
func (a action) cols() int {
	return a.c2 - a.c1 + 1
}

type ledGrid struct {
	*mat64.Dense
}

func (grid *ledGrid) apply(action action) {
	view := grid.View(action.r1, action.c1, action.rows(), action.cols()).(*mat64.Dense)
	var change func(int, int, float64) float64
	switch action.Verb {
	case toggle:
		change = puzzleToggle
	case turnOn:
		change = puzzleTurnOn
	case turnOff:
		change = puzzleTurnOff
	}
	view.Apply(change, view)
}

func puzzle1Toggle(_ int, _ int, v float64) float64 {
	if v == 1 {
		return 0
	}
	return 1
}

func puzzle1TurnOn(int, int, float64) float64 {
	return 1
}

func puzzle1TurnOff(int, int, float64) float64 {
	return 0
}

func puzzle2Toggle(_ int, _ int, v float64) float64 {
	return v + 2.
}

func puzzle2TurnOn(_ int, _ int, v float64) float64 {
	return v + 1.
}

func puzzle2TurnOff(_ int, _ int, v float64) float64 {
	return max(0., v-1.)
}

func max(a, b float64) float64 {
	if a < b {
		return b
	}
	return a
}

func switchToDialect1() {
	puzzleToggle = puzzle1Toggle
	puzzleTurnOff = puzzle1TurnOff
	puzzleTurnOn = puzzle1TurnOn
}

func switchToDialect2() {
	puzzleToggle = puzzle2Toggle
	puzzleTurnOff = puzzle2TurnOff
	puzzleTurnOn = puzzle2TurnOn
}
