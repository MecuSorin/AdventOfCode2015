/* Author: Mecu Sorin       Phone: 0040747020102 */

package Day07

import (
	"fmt"
	"strconv"
)

type lexonMatcherFn func(lexonType) bool

func receiver(typ lexonType) bool {
	return lexonVariable == typ
}

func provider(typ lexonType) bool {
	return lexonVariable == typ || lexonNumber == typ
}

func keyword(typ lexonType) lexonMatcherFn {
	return func(lexedType lexonType) bool {
		return typ == lexedType
	}
}

type parser struct {
	format []lexonMatcherFn
	action func(state map[string]int, lexons []lexon) error
}

func (l *lexer) collect() []lexon {
	result := []lexon{}
	for lxn := range l.emitter {
		if lexonNumber == lxn.typ {
			_, err := strconv.Atoi(lxn.val)
			if nil != err {
				lxn.typ = lexonError
			}
		}
		result = append(result, lxn)
	}
	return result
}

func validateNoErrors(lexons []lexon) error {
	for _, lxn := range lexons {
		if lexonError == lxn.typ {
			return fmt.Errorf("Lexer failed to interpret the entry at position %d %q\n", lxn.pos, lxn.val)
		}
	}
	return nil
}

func identifyParser(parsers []parser, lexons []lexon) (expression, error) {
	if err := validateNoErrors(lexons); nil != err {
		return expression{}, err
	}

	for _, p := range parsers {
		matched := true
		if len(p.format) != len(lexons) {
			continue
		}
		for i, lxn := range lexons {
			matched = matched && p.format[i](lxn.typ)
		}
		if matched {
			return newExpression(p, lexons), nil
		}
	}
	return expression{}, fmt.Errorf("Unknown expression grammar: %v", lexons)
}

func process(instructions []string) (map[string]int, error) {
	expressions := make([]expression, len(instructions))
	for i, t := range instructions {
		l := newLexer(t)
		lexons := l.collect()
		if 0 == len(lexons) {
			continue
		}
		statement, err := identifyParser(knownParsers, lexons)

		if nil != err {
			return nil, err
		}
		expressions[i] = statement
	}
	return evaluateExpressions(expressions)
}

var (
	parserSET = parser{
		[]lexonMatcherFn{provider, keyword(lexonKeywordEndGateTerminal), receiver},
		func(state map[string]int, lexons []lexon) error {
			v, err := getValue(lexons[0], state)
			if nil != err {
				return err
			}
			state[lexons[2].val] = v
			return nil
		},
	}
	parserNOT = parser{
		[]lexonMatcherFn{keyword(lexonKeywordNOT), provider, keyword(lexonKeywordEndGateTerminal), receiver},
		func(state map[string]int, lexons []lexon) error {
			v, err := getValue(lexons[1], state)
			if nil != err {
				return err
			}
			state[lexons[3].val] = int(^uint16(v))
			return nil
		},
	}

	parserAND    = generateParser(provider, lexonKeywordAND, provider, func(a, b int) int { return a & b })
	parserOR     = generateParser(provider, lexonKeywordOR, provider, func(a, b int) int { return a | b })
	parserLSHIFT = generateParser(provider, lexonKeywordLSHIFT, provider, func(a, b int) int { return a << uint(b) })
	parserRSHIFT = generateParser(provider, lexonKeywordRSHIFT, provider, func(a, b int) int { return a >> uint(b) })

	knownParsers = []parser{parserAND, parserNOT, parserLSHIFT, parserOR, parserRSHIFT, parserSET}
)

func getValue(l lexon, state map[string]int) (int, error) {
	switch l.typ {
	case lexonNumber:
		return strconv.Atoi(l.val)
	case lexonVariable:
		v, ok := state[l.val]
		if !ok {
			return 0, fmt.Errorf("Variable %q is not initialized\n", l.val)
		}
		return v, nil

	}
	return 0, fmt.Errorf("getValue doesn't support lexon %v", l)
}

func generateParser(first lexonMatcherFn, operatorKeyword lexonType, second lexonMatcherFn, operator func(a, b int) int) parser {
	return parser{
		[]lexonMatcherFn{first, keyword(operatorKeyword), second, keyword(lexonKeywordEndGateTerminal), receiver},
		func(state map[string]int, lexons []lexon) error {
			v1, err := getValue(lexons[0], state)
			if nil != err {
				return err
			}
			v2, err := getValue(lexons[2], state)
			if nil != err {
				return err
			}
			state[lexons[4].val] = operator(v1, v2)
			return nil
		},
	}
}
