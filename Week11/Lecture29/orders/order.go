package orders

type Order struct {
	Customer string
	Amount   int
}

func GroupBy[T any, U comparable](col []T, keyFn func(T) U) map[U][]T {
	m := make(map[U][]Order)

	for _, v := range col {
		if typeAssert, ok := any(v).(Order); ok {
			m[keyFn(v)] = append(m[keyFn(v)], Order{Amount: typeAssert.Amount})
		}
	}

	var orderToTtype map[U][]T = any(m).(map[U][]T)

	return orderToTtype
}
