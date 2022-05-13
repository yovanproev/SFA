package final

import (
	"testing"
)

func TestInitDB(t *testing.T) {
	filepath := ":memory:"
	got := OpenDBConnection(filepath)

	if got == nil {
		t.Fatalf("Expected database, got %+v", got)
	}
}

// go test . -v -cover
// === RUN   TestInitDB
// --- PASS: TestInitDB (0.00s)
// PASS
// coverage: 78.9% of statements
// ok      final/cmd/echo/DBInit   0.266s  coverage: 78.9% of statements
