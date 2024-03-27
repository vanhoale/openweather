package api

import (
	"encoding/json"
	"log"
	"net/http"
	"openweather/model"

	"github.com/go-resty/resty/v2"
)

func GetWeatherByLatLong(lat string, lon string, apiKey string) model.ApiResp {
	client := resty.New()
	openWeatherResp := model.OpenWeatherResp{}
	resp, err1 := client.R().
		SetQueryParams(map[string]string{
			"lat":   lat,
			"lon":   lon,
			"appid": apiKey,
			"units": "imperial",
		}).
		SetHeader("Accept", "application/json").
		SetHeader("Content-Type", "application/json").
		Get("https://api.openweathermap.org/data/2.5/weather")

	if err1 != nil || resp.StatusCode() != http.StatusOK {
		log.Fatal(err1.Error())
	}
	err2 := json.Unmarshal(resp.Body(), &openWeatherResp)
	if err2 != nil {
		log.Fatal(err2.Error())
	}
	feel := "Cold"
	if openWeatherResp.Main.FeelsLike > 24 && openWeatherResp.Main.FeelsLike < 86 {
		feel = "Moderate"
	} else if openWeatherResp.Main.FeelsLike >= 86 {
		feel = "Hot"
	}
	data := model.ApiResp{Weather: openWeatherResp.Weather[0].Main, Feel: feel}

	return data
}
