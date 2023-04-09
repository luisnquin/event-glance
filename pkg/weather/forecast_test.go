package weather_test

import (
	"fmt"
	"testing"

	"github.com/goccy/go-json"
	"github.com/luisnquin/event-glance/pkg/weather"
)

func TestSmoke(t *testing.T) {
	response, err := weather.Forecast(-9.125280, -78.528610, weather.WithCurrent())
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	data, err := json.Marshal(response)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", data)
}
