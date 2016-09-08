/* Author: Mecu Sorin       Phone: 0040747020102 */

package Day09

import (
	"regexp"
	"strconv"
)

var (
	roadRegexPattern = regexp.MustCompile(`(\w+)\s+to\s+(\w+)\s*=\s*(\d+)`)
)

type city string
type unorderedRoads map[city]uint
type unorderedRoadsLayout map[city]unorderedRoads

func newRoadsLayout() unorderedRoadsLayout {
	return unorderedRoadsLayout(make(map[city]unorderedRoads))
}

func (r unorderedRoadsLayout) addRoad(fromCity, toCity city, distance uint) {
	r.addCity(fromCity, toCity, distance)
	r.addCity(toCity, fromCity, distance)
}

func (r unorderedRoadsLayout) addCity(fromCity, toCity city, distance uint) {
	fromCityRoads, ok := r[fromCity]
	if !ok {
		fromCityRoads = unorderedRoads(make(map[city]uint))
		r[fromCity] = fromCityRoads
	}
	fromCityRoads[toCity] = distance
}

func readRoads(serializedRoads []string) unorderedRoadsLayout {
	cities := newRoadsLayout()
	for _, serializedRoad := range serializedRoads {
		m := roadRegexPattern.FindAllStringSubmatch(serializedRoad, -1)
		if nil == m {
			continue
		}
		distance, _ := strconv.Atoi(m[0][3])
		cities.addRoad(city(m[0][1]), city(m[0][2]), uint(distance))
	}
	return cities
}
