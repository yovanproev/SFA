package main

import (
	"fmt"
	"sort"
	"time"
)

type ByDate []time.Time

func (a ByDate) Len() int           { return len(a) }
func (a ByDate) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByDate) Less(i, j int) bool { return a[i].Before(a[j]) }

func sortDates(format string, dates ...string) ([]string, error) {
	var datesInTimeFormat []time.Time
	var sortedDates []string
	var mismatch error

	for i := 0; i < len(dates); i++ {
		theTime, err := time.Parse(format, dates[i])

		if err != nil {
			mismatch = err
		}
		datesInTimeFormat = append(datesInTimeFormat, theTime)
	}
	sort.Sort(ByDate(datesInTimeFormat))

	for _, v := range datesInTimeFormat {
		sortedDates = append(sortedDates, v.String())
	}

	return sortedDates, mismatch
}

func main() {
	var dates = []string{"Sep-14-2008", "Dec--2021", "Mar-18-2022", "Apr-01-2006"}
	var format = "Jan-02-2006"

	sortingDates, err := sortDates(format, dates...)
	if err != nil {
		fmt.Println("Could not parse time: ", err)
	} else {
		fmt.Println(sortingDates)
	}

}

// Input: var dates = []string{"Sep-14-2008", "Dec-03-2021", "Mar-18-2022", "Apr-01-2006"}
// Output:
// [2006-04-01 00:00:00 +0000 UTC 2008-09-14 00:00:00 +0000 UTC 2021-12-03 00:00:00 +0000 UTC 2022-03-18 00:00:00 +0000 UTC]

// Input with error: var dates = []string{"Sep-14-2008", "Dec--2021", "Mar-18-2022", "Apr-01-2006"}
// Could not parse time:  parsing time "Dec--2021" as "Jan-02-2006": cannot parse "-2021" as "02"
