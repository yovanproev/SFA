package weather

import (
	"encoding/json"
	handleErrors "final/pkg/app/errors"
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
	handleError := handleErrors.Error{}.SetErrors()
	want := mockWeatherInfo

	router := http.NewServeMux()
	router.Handle("/", mockFetchWeather(mockWeather))
	mockServer := httptest.NewServer(router)

	got := FetchWeather("12", "24", "", mockServer.URL, handleError)

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

// $ go test . -v -cover
// === RUN   TestFetchWeather
// --- PASS: TestFetchWeather (0.00s)
// PASS
// coverage: 73.7% of statements
// ok      final/pkg/app/weather   0.288s  coverage: 73.7% of statements
