/* Author: Mecu Sorin       Phone: 0040747020102 */

package Day11

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

func TestBootstaping(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Santa must change passwords Suite")
}

var _ = Describe("When Santa must change password by succesive incrementations", func() {
	Specify("Interpreting an invalid password should fail", func() {
		_, err := getBase26FromString("a1a")
		Expect(err).ShouldNot(Succeed())
	})

	DescribeTable("Interpreting a good password should work",
		func(from string, expected int) {
			pass, err := getBase26FromString(from)
			Expect(err).Should(Succeed())
			Expect(pass).To(Equal(base26(expected)))
		},
		Entry("a", "a", 0),
		Entry("z", "z", 25),
		Entry("az", "az", 0+25),
		Entry("ba", "ba", 26+0),
		Entry("bb", "bb", 26+1),
	)

	DescribeTable("Round trip a password should work",
		func(from string) {
			pass, err := getBase26FromString(from)
			Expect(err).Should(Succeed())
			actual := pass.String()
			secondPass, err := getBase26FromString(actual)
			Expect(err).Should(Succeed())
			Expect(pass).To(Equal(secondPass))
		},
		Entry("a", "a"),
		Entry("z", "z"),
		Entry("az", "z"),
		Entry("ba", "ba"),
		Entry("bb", "bb"),
		Entry("abbcegjk", "abbcegjk"),
	)

	DescribeTable("Should match the given samples",
		func(fromPassword, expected string) {
			pass, err := getBase26FromString(fromPassword)
			Expect(err).Should(Succeed())
			Expect(base26(int64(pass) + int64(1)).String()).To(Equal(expected))
		},
		Entry("xx", "xx", "xy"),
		Entry("xy", "xy", "xz"),
		Entry("xz", "xz", "ya"),
		Entry("ya", "ya", "yb"),
		Entry("bzz", "bzz", "caa"),
	)
})
