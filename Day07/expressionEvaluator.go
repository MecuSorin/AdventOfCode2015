/* Author: Mecu Sorin       Phone: 0040747020102 */

package Day07

import (
	"fmt"
)

type expression struct {
	parser
	lexons []lexon
	failed error
}

var errorUnprocessed = fmt.Errorf("Unprocessed")

func newExpression(p parser, lexons []lexon) expression {
	return expression{p, lexons, errorUnprocessed}
}

func (e *expression) process(state map[string]int) (now bool, done bool) {
	if e.failed == nil {
		return false, true
	}
	e.failed = e.parser.action(state, e.lexons)
	done = nil == e.failed
	return true && done, done
}

func evaluateExpressions(expressions []expression) (map[string]int, error) {
	state := make(map[string]int)
	processed, remaining := true, false
	for processed {
		processed, remaining = false, false
		for i := range expressions {
			now, done := expressions[i].process(state)
			processed = processed || now
			remaining = remaining || !done
		}
	}
	if remaining {
		unprocessed := []error{}
		for _, e := range expressions {
			if nil != e.failed {
				unprocessed = append(unprocessed, e.failed)
			}
		}
		return state, fmt.Errorf("Failed because: %+v", unprocessed)
	}
	return state, nil
}
