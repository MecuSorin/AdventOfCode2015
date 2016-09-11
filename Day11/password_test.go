/* Author: Mecu Sorin       Phone: 0040747020102 */

package Day11

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var _ = Describe("When Santa must change password by succesive incrementations", func() {

	DescribeTable("Santa's password should comply the Security-Elf rule about 3 consecutive letters",
		func(password string, expectedValidate3ConsecutiveLetters bool) {
			pass, err := getBase26FromString(password)
			Expect(err).Should(Succeed())
			p := santaPassword{pass}
			Expect(p.has3ConsecutiveLetters()).To(Equal(expectedValidate3ConsecutiveLetters))
		},
		Entry("bc", "bc", true),
		Entry("xyz", "xyz", true),
		Entry("cd", "cd", false),
		Entry("hijklmmn", "hijklmmn", true),
		Entry("abbceffg", "abbceffg", false),
	)

	DescribeTable("Santa's password should comply the Security-Elf rule about confusing letters",
		func(password string, expectedValidateConfusingLetters bool) {
			pass, err := getBase26FromString(password)
			Expect(err).Should(Succeed())
			p := santaPassword{pass}
			Expect(p.hasConfusingLetters()).To(Equal(expectedValidateConfusingLetters))
		},
		Entry("bc", "bc", false),
		Entry("xiyz", "xiyz", true),
		Entry("lcd", "lcd", true),
		Entry("cdo", "cdo", true),
		Entry("hijklmmn", "hijklmmn", true),
	)

	DescribeTable("Santa's password should comply the Security-Elf rule about 2 distinct consecutive letter pairs",
		func(password string, expectedValidateConfusingLetters bool) {
			pass, err := getBase26FromString(password)
			Expect(err).Should(Succeed())
			p := santaPassword{pass}
			Expect(p.hasTwoDifferentLetterPairs()).To(Equal(expectedValidateConfusingLetters))
		},
		Entry("bc", "bc", false),
		Entry("xiyyz", "xiyyz", true),
		Entry("bblcd", "bblcd", true),
		Entry("cdaao", "cdaao", false),
		Entry("abbceffg", "abbceffg", true),
		Entry("abbcegjk", "abbcegjk", false),
	)

	DescribeTable("Santa's password should comply puzzle 1 samples",
		func(from, to string) {
			actual, err := getNextPassword(from)
			Expect(err).Should(Succeed())
			Expect(actual).To(Equal(to))

		},
		Entry("abcdefgh", "abcdefgh", "abcdffaa"),
		XEntry("ghijklmn", "ghijklmn", "ghjaabcc"), // slow test
	)

	Specify("Just for coverage sake", func() {
		_, err := getNextPassword("1")
		Expect(err).ShouldNot(Succeed())
	})

	Specify("Given puzzle 1 input: hepxcrrq", func() {
		actual, err := getNextPassword("hepxcrrq")
		Expect(err).Should(Succeed())
		Expect(actual).To(Equal("hepxxyzz"))
	})
	Specify("Given puzzle 2 input: hepxxyzz", func() {
		actual, err := getNextPassword("hepxxyzz")
		Expect(err).Should(Succeed())
		Expect(actual).To(Equal("heqaabcc"))
	})
})
