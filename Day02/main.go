/* Author: Mecu Sorin       Phone: 0040747020102 */

package day02

import (
	"fmt"
	"regexp"
	"strconv"
)

type gift struct {
	w, h, l int
}

var giftRegexPattern = regexp.MustCompile(`(?mi)^(\d+)x(\d+)x(\d+)$`)

func getWrappingMaterials(giftSizes string) (paper, ribbon int, err error) {
	gifts, err := readGifts(giftSizes)
	if nil != err {
		return 0, 0, err
	}

	for _, g := range gifts {
		paper += g.getWrappingPaper()
		ribbon += g.getRibbonNeededForWrapping()
	}
	return paper, ribbon, nil
}

func readGifts(from string) ([]gift, error) {
	results := giftRegexPattern.FindAllStringSubmatch(from, -1)
	if nil == results {
		return nil, fmt.Errorf("Unable to parse the list of gift sizes")
	}
	gifts := make([]gift, len(results))
	for i, v := range results {
		w, _ := strconv.Atoi(v[1])
		h, _ := strconv.Atoi(v[2])
		l, _ := strconv.Atoi(v[3])
		gifts[i] = gift{w: w, h: h, l: l}
	}
	return gifts, nil
}

func (g gift) getWrappingPaper() int {
	s1 := g.w * g.h
	s2 := g.w * g.l
	s3 := g.h * g.l
	return 2*(s1+s2+s3) + min(s1, min(s2, s3))
}

func (g gift) getRibbonNeededForWrapping() int {
	r1 := g.w + g.h
	r2 := g.w + g.l
	r3 := g.h + g.l
	return 2*min(r1, min(r2, r3)) + g.w*g.h*g.l
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
