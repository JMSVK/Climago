package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

var good bool = false

func main() {
	var ciudad string
	println("De que cuidad eres")
	fmt.Scan(&ciudad)
	var geocode string = "https://geocoding-api.open-meteo.com/v1/search?name=" + ciudad + "&count=1"
	ans, err := http.Get(geocode)
	if err != nil {
		println("Caramba")
	}
	body, crash := io.ReadAll(ans.Body)

	if crash != nil {
		println("Caramba")
	}

	type Coordenadas struct {
		Latitud   float64 `json:"latitude"`
		Longitud  float64 `json:"longitude"`
		Elevation float64 `json:"elevation"`
	}

	type resultado struct {
		Results []Coordenadas `json:"results"`
	}

	var data resultado
	json.Unmarshal(body, &data)

	var lat string = fmt.Sprintf("%g", data.Results[0].Latitud)
	var long string = fmt.Sprintf("%g", data.Results[0].Longitud)

	var clima string = "https://api.open-meteo.com/v1/forecast?latitude=" + lat + "&longitude=" + long + "&current_weather=true"
	resp, unresp := http.Get(clima)

	if unresp != nil {
		println("Caramba")
	}

	cbody, ccrash := io.ReadAll(resp.Body)

	if ccrash != nil {
		println("Caramba")
	}

	if cbody != nil {
		good = true //:)
	}

	type Clima struct {
		Temperature float64 `json:"temperature"`
		Windspeed   float64 `json:"windspeed"`
		Is_day      int     `json:"is_day"`
	}

	type Climas struct {
		Current_weather Clima `json:"current_weather"`
	}

	var cdata Climas
	json.Unmarshal(cbody, &cdata)
	var ctemp = fmt.Sprintf("%v", cdata.Current_weather.Temperature)
	var cwinspeed = fmt.Sprintf("%v", cdata.Current_weather.Windspeed)
	var DoN string = "Nada"

	DoN = "Nada"

	if cdata.Current_weather.Is_day == 0 {
		DoN = "Noche"
	}

	if cdata.Current_weather.Is_day == 1 {
		DoN = "Dia"
	}

	println("La temperatura en " + ciudad + " es de " + ctemp + " grados celsius " + "la fuerza del viento es de " + cwinspeed + " km x h y es de " + DoN)
}
