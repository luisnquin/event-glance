package weather_test

import (
	"fmt"
	"testing"

	"github.com/goccy/go-json"
	"github.com/luisnquin/event-glance/pkg/weather"
)

func TestSmoke(t *testing.T) {
	response, err := weather.Forecast(-15.125280, 35.528610, weather.Temperature(weather.Celsius))
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	data, err := json.Marshal(response)
	if err != nil {
		panic(err)
	}

	fmt.Println(response.CurrentWeather.WeatherCode)
	fmt.Println(response.CurrentWeather)
	fmt.Printf("%s\n", data)
}
