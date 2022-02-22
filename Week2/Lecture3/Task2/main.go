package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	singleMap()
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

	return cities, prices
}

func singleMap() {
	cities, prices := citiesAndPrices()
	keys := make(map[string][]int)

	for i, city := range cities {
		keys[city] = prices[:4]
		i++
	}

	fmt.Println((keys))

}
