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
		keys[city] = valueSlice[:4]
		i++
	}
	fmt.Println(keys)
	return keys
}
