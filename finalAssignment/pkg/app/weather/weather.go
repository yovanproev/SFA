package weather

import (
	"encoding/json"
	handleErrors "final/pkg/app/errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
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

func FetchWeather(lat, lon string, apiKeys, url string, e handleErrors.Error) WeatherInfo {
	var weather Weather

	if apiKeys != "" {
		url = "https://api.openweathermap.org/data/2.5/weather?lat=" + lat + "&lon=" + lon + "&appid=" + apiKeys + "&units=metric"
	}

	client := http.Client{
		Timeout: time.Second * 2, // Timeout after 2 seconds
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatalf(e.HTTPRequest, err)
	}

	res, getErr := client.Do(req)
	if getErr != nil {
		log.Fatalf(e.HTTPResponse, getErr)
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatalf(e.HTTPResponse, readErr)
	}

	if err := json.Unmarshal(body, &weather); err != nil {
		log.Printf(e.JSONMarshalling, err)
	}

	wi := WeatherInfo{
		FormattedTemp: fmt.Sprintf("%f", weather.Main.Temp),
		Description:   weather.Weather[0].Description,
		City:          weather.Name,
	}

	return wi
}
