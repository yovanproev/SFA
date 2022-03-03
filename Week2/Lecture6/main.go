package main

import (
	"fmt"
	"sort"
)

type Item struct {
	Value    int
	PrevItem *Item
}

type MagicList struct {
	LastItem *Item
	length   int
}

func (l *MagicList) add(i *Item) {
	second := l.LastItem
	l.LastItem = i
	l.LastItem.PrevItem = second
	l.length++
}

func toSlice(ml MagicList) []int {
	var slice []int

	toPrint := ml.LastItem
	slice = append(slice, ml.LastItem.Value)

	for ml.length != 1 {
		toPrint = toPrint.PrevItem
		slice = append(slice, toPrint.Value)
		ml.length--
	}

	fmt.Println("original slice", slice)

	// reversing function
	for i := 0; i < len(slice)/2; i++ {
		j := len(slice) - i - 1
		slice[i], slice[j] = slice[j], slice[i]
	}
	fmt.Println("reverse the slice", slice)

	// sorting function, biggest number -> smallest
	sort.Sort(sort.Reverse(sort.IntSlice(slice)))

	fmt.Println("sort the slice, biggest->smallest", slice)

	return slice
}

func main() {
	ml := MagicList{}

	node1 := &Item{Value: 10}
	node2 := &Item{Value: 22}
	node3 := &Item{Value: 13}
	ml.add(node1)
	ml.add(node2)
	ml.add(node3)

	toSlice(ml)
}

// Output:
// original slice [13 22 10]
// reverse the slice [10 22 13]
// sort the slice, biggest->smallest [22 13 10].
