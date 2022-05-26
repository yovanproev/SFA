package config

import (
	"reflect"
	"testing"
)

func TestLoadEnv(t *testing.T) {
	configuration := Configurations{}
	configuration.SetConfig()

	want := LoadEnv(configuration.DevelopmentEnv, configuration)
	got := ""

	if !reflect.DeepEqual(got, want) {
		t.Fatalf(`
			Got %+v, 
			expected %+v`, got, want)
	}
}

// $ go test . -v -cover
// === RUN   TestLoadEnv
// --- PASS: TestLoadEnv (0.00s)
// PASS
// coverage: 90.9% of statements
// ok      final/pkg/config        0.221s  coverage: 90.9% of statements
