/* Author: Mecu Sorin       Phone: 0040747020102 */

package Day09

import (
	"bufio"
	"os"
	"runtime"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	//. "github.com/onsi/ginkgo/extensions/table"
)

func TestBootstaping(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, " Suite")
}

func loadCitiesFromFile(file string) unorderedRoadsLayout {
	roadsFile, err := os.Open(file)
	Expect(err).Should(Succeed())
	defer roadsFile.Close()
	scanner := bufio.NewScanner(roadsFile)
	roads := []string{}
	for scanner.Scan() {
		roads = append(roads, scanner.Text())
	}
	Expect(scanner.Err()).Should(Succeed())
	return readRoads(roads)
}

var _ = Describe("Santa's travel", func() {
	BeforeEach(func() {
		isABetterSolutionBasedOnDistance = isMax
	})

	Specify("Should know how many available procs to use", func() {
		oldMaxProcs := runtime.GOMAXPROCS(0)
		defer runtime.GOMAXPROCS(oldMaxProcs)
		runtime.GOMAXPROCS(1)
		Expect(maxParallelism()).To(Equal(1))
	})

	Specify("Can load cities from sample.input", func() {
		cities := loadCitiesFromFile("sample.input")
		Expect(len(cities)).To(Equal(3))
	})

	Specify("Should not fail if only two cities are provided", func() {
		const expectedDistance = uint(12)
		roads := unorderedRoadsLayout(make(map[city]unorderedRoads))
		roads["start"], roads["end"] = unorderedRoads(make(map[city]uint)), unorderedRoads(make(map[city]uint))
		roads["start"]["end"], roads["end"]["start"] = expectedDistance, expectedDistance
		shortestDistance, ok := findShortestDistance(roads)
		Expect(ok).To(BeTrue())
		Expect(len(shortestDistance.visited)).To(Equal(2))
		Expect(shortestDistance.distance).To(Equal(expectedDistance))
	})

	Specify("Sample provided from sample.input should work", func() {
		cities := loadCitiesFromFile("sample.input")
		solution, ok := findShortestDistance(cities)
		Expect(ok).To(BeTrue())
		Expect(len(solution.visited)).To(Equal(3))
		Expect(solution.distance).To(Equal(uint(141 + 464)))

	})
})
