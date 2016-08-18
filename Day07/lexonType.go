/* Author: Mecu Sorin       Phone: 0040747020102 */

package Day07

type lexonType int

type lexon struct {
	typ lexonType
	pos Pos
	val string
}

const (
	lexonError lexonType = iota
	lexonKeywordAND
	lexonKeywordEndGateTerminal
	lexonKeywordLSHIFT
	lexonKeywordNOT
	lexonKeywordOR
	lexonKeywordRSHIFT
	lexonNumber
	lexonVariable

	gateEndTerminal = "->"
)

var grammarKeywords = map[string]lexonType{
	"AND":    lexonKeywordAND,
	"OR":     lexonKeywordOR,
	"LSHIFT": lexonKeywordLSHIFT,
	"RSHIFT": lexonKeywordRSHIFT,
	"NOT":    lexonKeywordNOT,
}
