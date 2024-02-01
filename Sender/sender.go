package sender

import data "github.com/karalakrepp/AutomatedWeatherPolling/Data"

type Sender interface {
	Send(*data.WeatherData) error
}
