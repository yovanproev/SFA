package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	citiesAndPrices()
}

func citiesAndPrices() ([]string, []int) {
	rand.Seed(time.Now().UnixMilli())
	cityChoices := []string{"Berlin", "Moscow", "Chicago", "Tokyo", "London"}
	dataPointCount := 100

	//randomly choise cities
	cities := make([]string, dataPointCount)
	for i := range cities {
		cities[i] = cityChoices[rand.Intn(len(cityChoices))]
	}
	prices := make([]int, dataPointCount)
	for i := range prices {
		prices[i] = rand.Intn(100)
	}
	groupSlice(cities, prices)
	return cities, prices
}

func groupSlice(keySlice []string, valueSlice []int) map[string][]int {
	keys := make(map[string][]int)

	for i, city := range keySlice {
		var limitingSlice int
		if i < valueSlice[i] {
			limitingSlice = valueSlice[i] - i
			keys[city] = valueSlice[i : valueSlice[i]-limitingSlice+4]
		} else {
			limitingSlice = i - valueSlice[i]
			keys[city] = valueSlice[valueSlice[i] : i-limitingSlice+4]
		}
		i++
	}
	fmt.Print(keys)
	return keys
}

//Output:
// map[Berlin:[96 65 56 30] Chicago:[29 41 98 94] London:[89 35 75 23] Moscow:[29 41 98 94] Tokyo:[16 49 3 73]]
// map[Berlin:[21 35 9 62] Chicago:[87 36 37 27] London:[84 69 22 67] Moscow:[20 56 68 60] Tokyo:[27 5 15 91]]
// map[Berlin:[41 62 27 78] Chicago:[80 25 15 73] London:[5 90 71 59] Moscow:[90 71 59 63] Tokyo:[15 9 88 72]]
