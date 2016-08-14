/* Author: Mecu Sorin       Phone: 0040747020102 */

package Day05

import (
	"bufio"
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Santa categorizing letters", func() {
	tester := func(categorizer func(string) bool, expectedNiceCount int) {
		samples, err := os.Open("puzzle1.txt")
		Expect(err).Should(Succeed())
		defer samples.Close()
		scanner := bufio.NewScanner(samples)
		counter := 0
		for scanner.Scan() {
			if categorizer(scanner.Text()) {
				counter++
			}
		}
		Expect(scanner.Err()).Should(Succeed())
		Expect(counter).To(Equal(expectedNiceCount))
	}

	Specify("First puzzle", func() {
		tester(categorize, 255)
	})
	Specify("Second puzzle", func() {
		tester(categorizePuzzle2, 55)
	})
})
