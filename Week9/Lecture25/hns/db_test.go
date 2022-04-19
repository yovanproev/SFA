package hns

import (
	"testing"
	"time"
)

func TestInitDB(t *testing.T) {
	filepath := ":memory:"
	db := InitDB(filepath)

	if db == nil {
		t.Fatal("Database file not present, database not initialized!")
	}
}

func TestCreateAndStoreAndReadTable(t *testing.T) {
	filepath := ":memory:"
	db := InitDB(filepath)
	if db == nil {
		t.Fatal("Database file not present, database not initialized!")
	}

	CreateTable(db)

	var ts TopStories
	stories := []Story{
		{Title: "First Title", Score: 1, DateStamp: time.Now()},
		{Title: "Second Title", Score: 2, DateStamp: time.Now()},
	}

	StoreItem(db, stories)

	readItems := ts.ReadItem(db)
	if readItems.Story == nil {
		t.Fatal("No stories to show!")
	}
}
