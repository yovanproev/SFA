package db

import (
	handleErrors "final/pkg/app/errors"
	"testing"
)

func TestInitDB(t *testing.T) {
	handleError := handleErrors.Error{}.SetErrors()

	filepath := ":memory:"
	got := OpenDBConnection(filepath, handleError)

	if got == nil {
		t.Fatalf("Expected database, got %+v", got)
	}
}

// $ go test . -v -covr
// === RUN   TestInitDB
// --- PASS: TestInitDB (0.00s)
// PASS
// coverage: 79.2% of statements
// ok      final/cmd/echo/DBInit   0.322s  coverage: 79.2% of statements
