package benchmarking

import (
	"testing"
	"time"
)

var primeNumberCap = 100

func Benchmark100PrimesWith0MSSleep(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PrimesAndSleep(primeNumberCap, 0*time.Millisecond)
	}
}

func Benchmark100PrimesWith5MSSleep(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PrimesAndSleep(primeNumberCap, 5*time.Millisecond)
	}
}

func Benchmark100PrimesWith10MSSleep(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PrimesAndSleep(primeNumberCap, 10*time.Millisecond)
	}
}

func Benchmark100GoPrimesWith0MSSleep(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GoPrimesAndSleep(primeNumberCap, 0*time.Millisecond)
	}
}

func Benchmark100GoPrimesWith5MSSleep(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GoPrimesAndSleep(primeNumberCap, 5*time.Millisecond)
	}
}

func Benchmark100GoPrimesWith10MSSleep(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GoPrimesAndSleep(primeNumberCap, 10*time.Millisecond)
	}
}

// Output
// $ go test -bench=.
// goos: windows
// goarch: amd64
// pkg: benchmarking/benchmarking
// cpu: Intel(R) Core(TM) i5-4210U CPU @ 1.70GHz
// Benchmark100PrimesWith0MSSleep-4          825924              1376 ns/op
// Benchmark100PrimesWith5MSSleep-4              10         106908590 ns/op
// Benchmark100PrimesWith10MSSleep-4             10         108424850 ns/op
// Benchmark100GoPrimesWith0MSSleep-4          1405            845685 ns/op
// Benchmark100GoPrimesWith5MSSleep-4          1528            840339 ns/op
// Benchmark100GoPrimesWith10MSSleep-4         1750           1108349 ns/op
// PASS
// ok      benchmarking/benchmarking       14.467s

// Output
// $ go test -bench=. -benchmem
// goos: windows
// goarch: amd64
// pkg: benchmarking/benchmarking
// cpu: Intel(R) Core(TM) i5-4210U CPU @ 1.70GHz
// Benchmark100PrimesWith0MSSleep-4          863491              1592 ns/op             120 B/op          4 allocs/op
// Benchmark100PrimesWith5MSSleep-4              10         107950360 ns/op             129 B/op          4 allocs/op
// Benchmark100PrimesWith10MSSleep-4             10         108102480 ns/op             129 B/op          4 allocs/op
// Benchmark100GoPrimesWith0MSSleep-4          1405            844460 ns/op           61985 B/op        490 allocs/op
// Benchmark100GoPrimesWith5MSSleep-4          1579            849689 ns/op           61861 B/op        497 allocs/op
// Benchmark100GoPrimesWith10MSSleep-4         1622           2435295 ns/op           63929 B/op        497 allocs/op
// PASS
// ok      benchmarking/benchmarking       15.762s
