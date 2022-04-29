package main

import "orders/Lecture29/orders"

func main() {
	Orders := []orders.Order{
		{Customer: "John", Amount: 1000},
		{Customer: "Sara", Amount: 2000},
		{Customer: "Sara", Amount: 1800},
		{Customer: "John", Amount: 1200},
	}

	orders.GroupBy(Orders, func(o orders.Order) string { return o.Customer })
}
