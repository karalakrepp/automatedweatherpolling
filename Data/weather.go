package data

type WeatherData struct {
	Elevation float64                `json :"elevation"`
	Hourly    map[string]interface{} `json :"elevation"`
}
