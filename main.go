package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"openweather/api"
	"os"
)

var OPENWEATHER_API_KEY = os.Getenv("OPENWEATHER_API_KEY")

func handler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	lat := r.Form.Get("lat")
	lon := r.Form.Get("lon")
	fmt.Println("Latitude: " + lat)
	fmt.Println("Longitude: " + lon)
	owResp := api.GetWeatherByLatLong(lat, lon, OPENWEATHER_API_KEY)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(owResp)

}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/weather", handler)
	fmt.Fprintf(os.Stdout, "Http Server started. Listening on 0.0.0.0:8080\n")
	http.ListenAndServe(":8080", mux)
}
