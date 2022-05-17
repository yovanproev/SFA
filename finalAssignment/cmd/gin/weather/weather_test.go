package final

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

var (
	mockWeatherInfo = WeatherInfo{
		FormattedTemp: "32.190000",
		Description:   "clear sky",
		City:          "Turabah",
	}

	mockWeather = Weather{
		Weather: []struct {
			ID          int    "json:\"id\""
			Main        string "json:\"main\""
			Description string "json:\"description\""
			Icon        string "json:\"icon\""
		}{{
			Description: "clear sky",
		}},
		Main: struct {
			Temp      float64 "json:\"temp\""
			FeelsLike float64 "json:\"feels_like\""
			TempMin   float64 "json:\"temp_min\""
			TempMax   float64 "json:\"temp_max\""
			Pressure  int     "json:\"pressure\""
			Humidity  int     "json:\"humidity\""
			SeaLevel  int     "json:\"sea_level\""
			GrndLevel int     "json:\"grnd_level\""
		}{
			Temp: 32.190000,
		},
		Name: "Turabah",
	}
)

func TestFetchWeather(t *testing.T) {
	want := mockWeatherInfo

	router := http.NewServeMux()
	router.Handle("/", mockFetchWeather(mockWeather))
	mockServer := httptest.NewServer(router)

	got := FetchWeather(12, 24, "", mockServer.URL)

	if !reflect.DeepEqual(got, want) {
		t.Fatalf(`
			Got %+v, 
			expected %+v`, got, want)
	}
}

func mockFetchWeather(wi Weather) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(wi)
	}
}

func TestLoadEnv(t *testing.T) {
	want := LoadEnv("development.env")
	got := ""

	if !reflect.DeepEqual(got, want) {
		t.Fatalf(`
			Got %+v, 
			expected %+v`, got, want)
	}
}

// $ go test . -v -cover
// === RUN   TestFetchWeather
// --- PASS: TestFetchWeather (0.00s)
// === RUN   TestLoadEnv
// --- PASS: TestLoadEnv (0.00s)
// PASS
// coverage: 81.0% of statements
// ok      final/cmd/echo/weather  (cached)        coverage: 81.0% of statements
