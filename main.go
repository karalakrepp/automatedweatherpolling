package main

import (
	sender "github.com/karalakrepp/AutomatedWeatherPolling/Sender"
	"github.com/karalakrepp/AutomatedWeatherPolling/handle"
)

func main() {

	SMSsender := sender.NewSMSSender("55555555")
	EmailSender := sender.NewEmailSender("xxxxxxxxxxp@gmail.com")

	wp := handle.NewWPoller(SMSsender, EmailSender)

	wp.Start()

}
