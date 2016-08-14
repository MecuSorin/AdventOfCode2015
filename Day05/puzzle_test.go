/* Author: Mecu Sorin       Phone: 0040747020102 */

package Day05

import (
	"bufio"
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Santa categorizing letters", func() {
	Specify("First puzzle", func() {
		samples, err := os.Open("puzzle1.txt")
		Expect(err).Should(Succeed())
		defer samples.Close()
		scanner := bufio.NewScanner(samples)
		counter := 0
		for scanner.Scan() {
			if categorize(scanner.Text()) {
				counter++
			}
		}
		Expect(scanner.Err()).Should(Succeed())
		Expect(counter).To(Equal(255))
	})
})
