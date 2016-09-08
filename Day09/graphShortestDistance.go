/* Author: Mecu Sorin       Phone: 0040747020102 */

package Day09

import (
	"runtime"
	"sync"
)

type toCity struct {
	toCity   city
	distance uint
}

type visitedCity struct {
	fromCity city
	roadUsed int
}

type shortestDistanceJob struct {
	cities       map[city][]toCity
	startingCity city
}

type shortestDistanceJobResult struct {
	distance uint
	visited  []city
}

func findShortestDistance(roadsLayout unorderedRoadsLayout) (shortestDistanceJobResult, bool) {
	jobs := make(chan shortestDistanceJob)
	solutionsProvidedByWorker := make(chan shortestDistanceJobResult)
	go func() {
		defer close(jobs)
		cities := getOrderedCitiesRoads(roadsLayout)
		for k := range roadsLayout {
			jobs <- shortestDistanceJob{cities, k}
		}
	}()
	go func() {
		defer close(solutionsProvidedByWorker)
		var ender sync.WaitGroup
		for worker := 0; worker < maxParallelism(); worker++ {
			ender.Add(1)
			go findShortestDistanceBacktracking(jobs, solutionsProvidedByWorker, &ender)
		}
		ender.Wait()
	}()
	solutionFound := false
	minDistance := shortestDistanceJobResult{}
	for result := range solutionsProvidedByWorker {
		if !solutionFound || minDistance.distance > result.distance {
			solutionFound = true
			minDistance = result
		}
	}
	return minDistance, solutionFound
}

func maxParallelism() int {
	maxProcs := runtime.GOMAXPROCS(0)
	numCPU := runtime.NumCPU()
	if maxProcs < numCPU && numCPU > 1 {
		return maxProcs
	}
	return numCPU
}

func getShortestDistanceJobResult(visitations []visitedCity, cities map[city][]toCity) shortestDistanceJobResult {
	visitationsLength := len(visitations) - 1
	visited := make([]city, visitationsLength+1)
	distance := 0
	for i := 0; i < visitationsLength; i++ {
		visited[i] = visitations[i].fromCity
		distance += int(cities[visited[i]][visitations[i].roadUsed-1].distance)
	}
	visited[visitationsLength] = visitations[visitationsLength].fromCity
	return shortestDistanceJobResult{uint(distance), visited}
}

func alreadyVisited(city city, visitations []visitedCity) bool {
	for i := range visitations {
		if visitations[i].fromCity == city {
			return true
		}
	}
	return false
}

func tryNextRoad(visitations []visitedCity, currentVisitOrder int, roadsFromCurrentCity []toCity) bool {
	for len(roadsFromCurrentCity) > visitations[currentVisitOrder].roadUsed {
		visitations[currentVisitOrder].roadUsed++
		toCity := roadsFromCurrentCity[visitations[currentVisitOrder].roadUsed-1].toCity
		if !alreadyVisited(toCity, visitations[:1+currentVisitOrder]) {
			visitations[1+currentVisitOrder] = visitedCity{toCity, 0}
			return true
		}
	}
	return false
}

func findShortestDistanceBacktracking(jobs <-chan shortestDistanceJob,
	result chan<- shortestDistanceJobResult,
	ender *sync.WaitGroup) {
	for job := range jobs {
		citiesLength := len(job.cities)
		visitations := make([]visitedCity, citiesLength)
		currentVisitOrder := 0
		visitations[0] = visitedCity{job.startingCity, 0}

		for currentVisitOrder > -1 {
			if currentVisitOrder >= citiesLength-1 {
				result <- getShortestDistanceJobResult(visitations, job.cities)
				currentVisitOrder--
				continue
			}
			currentCity := visitations[currentVisitOrder].fromCity
			if tryNextRoad(visitations, currentVisitOrder, job.cities[currentCity]) {
				currentVisitOrder++
				continue
			}
			currentVisitOrder--
		}
	}
	ender.Done()
}

func getOrderedCitiesRoads(roadsLayout unorderedRoadsLayout) map[city][]toCity {
	result := make(map[city][]toCity, len(roadsLayout))
	for c := range roadsLayout {
		result[c] = getOrderedDistances(c, roadsLayout[c])
	}
	return result
}

func getOrderedDistances(city city, roads unorderedRoads) []toCity {
	distances := make([]toCity, len(roads))
	cityIndex := 0
	for c, distance := range roads {
		distances[cityIndex] = toCity{c, distance}
		cityIndex++
	}
	return distances
}

type deliveredLayout map[city]uint
type trip struct {
	visited          deliveredLayout
	currentCityOrder uint
}
