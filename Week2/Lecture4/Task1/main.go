package main

import (
	"fmt"
	"time"
)

func main() {
	daysInMonth(2, 2020)
}

func daysInMonth(month, year int) int {
	months := time.Month(month)
	numberOfDays := time.Date(year, months+1, 0, 0, 0, 0, 0, time.UTC).Day()

	if month > 12 || month < 1 {
		numberOfDays = 0
	}

	fmt.Println(numberOfDays)

	return numberOfDays
}
