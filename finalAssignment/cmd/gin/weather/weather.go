package final

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type Weather struct {
	Coord struct {
		Lon float32 `json:"lon"`
		Lat float32 `json:"lat"`
	} `json:"coord"`
	Weather []struct {
		ID          int    `json:"id"`
		Main        string `json:"main"`
		Description string `json:"description"`
		Icon        string `json:"icon"`
	} `json:"weather"`
	Base string `json:"base"`
	Main struct {
		Temp      float64 `json:"temp"`
		FeelsLike float64 `json:"feels_like"`
		TempMin   float64 `json:"temp_min"`
		TempMax   float64 `json:"temp_max"`
		Pressure  int     `json:"pressure"`
		Humidity  int     `json:"humidity"`
		SeaLevel  int     `json:"sea_level"`
		GrndLevel int     `json:"grnd_level"`
	} `json:"main"`
	Visibility int `json:"visibility"`
	Wind       struct {
		Speed float64 `json:"speed"`
		Deg   int     `json:"deg"`
		Gust  float64 `json:"gust"`
	} `json:"wind"`
	Clouds struct {
		All int `json:"all"`
	} `json:"clouds"`
	Dt  int `json:"dt"`
	Sys struct {
		Country string `json:"country"`
		Sunrise int    `json:"sunrise"`
		Sunset  int    `json:"sunset"`
	} `json:"sys"`
	Timezone int    `json:"timezone"`
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Cod      int    `json:"cod"`
}

type WeatherInfo struct {
	FormattedTemp string `json:"formatted_temp,omitempty"`
	Description   string `json:"description,omitempty"`
	City          string `json:"city,omitempty"`
}

func LoadEnv(s string) string {
	err := godotenv.Load(s)
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}
	keys := os.Getenv("WEATHER_API_KEY")

	return keys
}

func FetchWeather(lat, lon float32, apiKeys, url string) WeatherInfo {
	var weather Weather

	latitude := fmt.Sprintf("%f", lat)
	longitude := fmt.Sprintf("%f", lon)

	var resp *http.Response
	var err error

	if apiKeys != "" {
		resp, err = http.Get("https://api.openweathermap.org/data/2.5/weather?lat=" + latitude + "&lon=" + longitude + "&appid=" + apiKeys + "&units=metric")
	} else {
		resp, err = http.Get(url)
	}
	if err != nil {
		log.Println("No response from request ", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err := json.Unmarshal(body, &weather); err != nil {
		log.Printf("Cannot unmarshal %s", err)
	}

	wi := WeatherInfo{
		FormattedTemp: fmt.Sprintf("%f", weather.Main.Temp),
		Description:   weather.Weather[0].Description,
		City:          weather.Name,
	}

	return wi
}
