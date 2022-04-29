package hns

import (
	"testing"
)

func TestInitDB(t *testing.T) {
	filepath := ":memory:"
	db := InitDB(filepath)

	if db == nil {
		t.Fatal("Database file not present, database not initialized!")
	}

	CreateTable(db)
}
