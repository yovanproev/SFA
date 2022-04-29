package orders

import (
	"reflect"
	"testing"
)

func TestXxx(t *testing.T) {
	want := map[string][]Order{
		"John": {
			{Amount: 1000},
			{Amount: 1200},
		},
		"Sara": {
			{Amount: 2000},
			{Amount: 1800},
		},
	}

	got := GroupBy([]Order{
		{Customer: "John", Amount: 1000},
		{Customer: "Sara", Amount: 2000},
		{Customer: "Sara", Amount: 1800},
		{Customer: "John", Amount: 1200},
	}, func(o Order) string {
		return o.Customer
	})

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("Expected %+v, got %+v", want, got)
	}
}
