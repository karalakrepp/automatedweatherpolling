package handle

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	data "github.com/karalakrepp/AutomatedWeatherPolling/Data"
	sender "github.com/karalakrepp/AutomatedWeatherPolling/Sender"
)

var (
	poolInterval = time.Second * 2
)

const (
	endpoint = "https://api.open-meteo.com/v1/forecast" //?latitude=52.52&longitude=13.41&current=temperature_2m&hourly=temperature_2m"

)

type WPoller struct {
	closech chan struct{}
	senders []sender.Sender
}

func NewWPoller(sender ...sender.Sender) *WPoller {
	return &WPoller{
		closech: make(chan struct{}),
		senders: sender,
	}
}

func (wp *WPoller) Start() {

	fmt.Println("WPOLLOR is starting")

	ticker := time.NewTicker(poolInterval)
free:
	for {
		select {

		case <-ticker.C:

			data, err := getWeatherResults(52.8, 13.6)

			if err != nil {
				log.Fatal(err)
			}

			if err := wp.handleData(data); err != nil {
				log.Fatal(err)
			}
		case <-wp.closech:
			//gracefull shutdown
			break free

		}

	}
	fmt.Println("wpoller stopped gracefully")
}

func getWeatherResults(lat, long float64) (*data.WeatherData, error) {

	uri := fmt.Sprintf("%s?latitude=%.2f&longitude=%.2f&current=temperature_2m&hourly=temperature_2m", endpoint, lat, long)

	fmt.Println("---------------------------------------------------------------------------------------------------------")

	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		return nil, err
	}

	client := http.Client{}

	res, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	var data data.WeatherData
	if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
		return nil, err
	}

	return &data, nil
}

func (wp *WPoller) handleData(data *data.WeatherData) error {

	for _, sender := range wp.senders {

		sender.Send(data)
	}

	return nil

}

// Close channel for gracefull shutdown
func (wp *WPoller) close() {
	close(wp.closech)
}
