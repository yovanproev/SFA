package hns

import (
	"reflect"
	"testing"
)

func TestInitializeDB(t *testing.T) {
	var want = SummedType{}

	db, got := InitializeDB(":memory:")

	if db == nil {
		t.Fatal("No database present!")
	}

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("Expected %+v, got %+v", want, got)
	}

}

// $ go test . -v -cover
// === RUN   TestInitDB
// --- PASS: TestInitDB (0.00s)
// === RUN   TestInitializeDB
// --- PASS: TestInitializeDB (0.00s)
// PASS
// coverage: 57.7% of statements
// ok      hns/dbInit      0.185s  coverage: 57.7% of statements
