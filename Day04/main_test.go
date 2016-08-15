/* Author: Mecu Sorin       Phone: 0040747020102 */

package day04

import (
	"runtime"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

func TestBootstaping(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Santa is mining AdventCoins Suite")
}

var _ = Describe("Santa is searching for numbers to generate special md5 hashes:", func() {
	Specify("Should know how many available procs to use", func() {
		oldMaxProcs := runtime.GOMAXPROCS(0)
		defer runtime.GOMAXPROCS(oldMaxProcs)
		runtime.GOMAXPROCS(1)
		Expect(maxParallelism()).To(Equal(1))
	})

	Specify("Should return error when an ill formated regex pattern is provided", func() {
		_, err := findNumberThatGenerateAMD5HashWithPatternSynchronous("aaa", "(", 0)
		Expect(err).ShouldNot(Succeed())
		_, err = findNumberThatGenerateAMD5HashWithPatternAsynchronous("aaa", "(", 0)
		Expect(err).ShouldNot(Succeed())
	})
	DescribeTable("Asynchronously searching md5 hashes that start with 00000",
		func(secret string, expected int) {
			Expect(findNumberThatGenerateAMD5HashWithPatternAsynchronous(secret, "^0{5}.*", 0)).To(Equal(expected))
		},
		Entry("MD5 Sample 01", "abcdef", 609043),
		Entry("MD5 Sample 01", "pqrstuv", 1048970),
	)

	DescribeTable("Synchronously searching md5 hashes that start with 00000",
		func(secret string, expected int) {
			Expect(findNumberThatGenerateAMD5HashWithPatternSynchronous(secret, "^0{5}.*", 0)).To(Equal(expected))
		},
		Entry("MD5 Sample 01", "abcdef", 609043),
		Entry("MD5 Sample 01", "pqrstuv", 1048970),
	)

	DescribeTable("Puzzle questions",
		func(startingFrom int, secret, pattern string, expected int) {
			Expect(findNumberThatGenerateAMD5HashWithPatternSynchronous(secret, pattern, startingFrom)).To(Equal(expected))
		},
		Entry("First : start with 00000", 0, "ckczppom", "^0{5}.*", 117946),
		Entry("Second : start with 00000", 117945, "ckczppom", "^0{6}.*", 3938038),
	)

	Specify("Test md5 hashing method", func() {
		sample := "abcdef609043"
		sampleHash := "000001dbbfa"

		hexaHash := getHexadecimalMD5Hash(sample)
		Expect(hexaHash[:len(sampleHash)]).To(Equal(sampleHash))
	})
})
