package weather

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-resty/resty/v2"
)

// curl -s "https://api.open-meteo.com/v1/forecast?latitude=52.52&longitude=13.41&current_weather=true&hourly=temperature_2m,relativehumidity_2m,windspeed_10m"

type ForecastResponse struct {
	Latitude              float64        `json:"latitude"`
	Longitude             float64        `json:"longitude"`
	GenerationTimeMs      float64        `json:"generationtime_ms"`
	UTCOffsetSeconds      uint64         `json:"utc_offset_seconds"`
	Timezone              string         `json:"timezone"`
	Timezone_abbreviation string         `json:"timezone_abbreviation"`
	Elevation             float64        `json:"elevation"`
	CurrentWeather        CurrentWeather `json:"current_weather"`
	HourlyUnits           HourlyUnits    `json:"hourly_units"`
	Hourly                Hourly         `json:"hourly"`
}

func (f *ForecastResponse) In() {
}

type (
	ForecastOption func(*forecastOptions)

	forecastOptions struct {
		CurrentWeather bool
	}
)

func WithCurrent() ForecastOption {
	return func(fo *forecastOptions) {
		fo.CurrentWeather = true
	}
}

func Forecast(latitude, longitude float64, options ...ForecastOption) (*ForecastResponse, error) { // latitude=52.52&longitude=13.41
	var queryOpts forecastOptions

	for _, option := range options {
		option(&queryOpts)
	}

	hourlyOptions := []string{"temperature_2m", "relativehumidity_2m", "windspeed_10m"}

	res, err := resty.New().R().SetQueryParams(map[string]string{
		"current_weather": strconv.FormatBool(queryOpts.CurrentWeather),
		"hourly":          strings.Join(hourlyOptions, ","),
		"latitude":        fmt.Sprintf("%.f", latitude),
		"longitude":       fmt.Sprintf("%.f", longitude),
	}).Get("https://api.open-meteo.com/v1/forecast")
	if err != nil {
		return nil, err
	}

	if res.StatusCode() != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d, body: %s", res.StatusCode(), res.Body())
	}

	var forecast ForecastResponse

	if err := json.Unmarshal(res.Body(), &forecast); err != nil {
		return nil, err
	}

	return &forecast, nil
}
