package weather_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/luisnquin/event-glance/pkg/weather"
)

func TestSmoke(t *testing.T) {
	response, err := weather.Forecast(-15.125280, 35.528610, weather.WithCurrent())
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	data, err := response.AfterCurrentWeather(time.Hour)
	if err != nil {
		panic(err)
	}

	fmt.Println(response.CurrentWeather)
	fmt.Println(data)
}
