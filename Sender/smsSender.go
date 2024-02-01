package sender

import (
	"fmt"

	data "github.com/karalakrepp/AutomatedWeatherPolling/Data"
)

type SMSSender struct {
	number string
}

func NewSMSSender(number string) *SMSSender {
	return &SMSSender{
		number: number,
	}
}

func (s *SMSSender) Send(data *data.WeatherData) error {
	fmt.Println("sms sending to :", s.number)

	return nil

}
