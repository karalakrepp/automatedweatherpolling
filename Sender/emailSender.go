package sender

import (
	"fmt"

	data "github.com/karalakrepp/AutomatedWeatherPolling/Data"
)

type EmailSender struct {
	email string
}

func NewEmailSender(email string) *EmailSender {
	return &EmailSender{
		email: email,
	}
}

func (s *EmailSender) Send(data *data.WeatherData) error {
	fmt.Println("email sending to :", s.email)

	return nil

}
