package main

import "fmt"

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

func (l MagicList) prinListData() {
	toPrint := l.LastItem
	for l.length != 0 {
		fmt.Printf("%d ", toPrint.Value)
		toPrint = toPrint.PrevItem
		l.length--
	}
}

func main() {
	ml := MagicList{}

	node1 := &Item{Value: 10}
	node2 := &Item{Value: 22}
	node3 := &Item{Value: 13}
	ml.add(node3)
	ml.add(node2)
	ml.add(node1)
	ml.prinListData()
}
