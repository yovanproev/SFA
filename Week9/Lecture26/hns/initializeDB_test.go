package hns

import (
	"reflect"
	"testing"
)

func TestInitializeDB(t *testing.T) {
	want := TopStories{}

	db, got := InitializeDB(":memory:")

	if db == nil {
		t.Fatal("No database present!")
	}
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("Expected %+v, got %+v", want, got)
	}
}
