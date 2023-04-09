package weather

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

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

func (f *ForecastResponse) AfterCurrentWeather(d time.Duration) (HourlyData, error) {
	targetHour := f.CurrentWeather.Time.Add(d).Round(time.Hour)

	for index, t := range f.Hourly.Time {
		if t.Equal(targetHour) {
			return HourlyData{
				RelativeHumidity2M: f.Hourly.RelativeHumidity2M[index],
				Temperature2M:      f.Hourly.Temperature2M[index],
				WindSpeed10M:       f.Hourly.WindSpeed10M[index],
				Time:               f.Hourly.Time[index],
			}, nil
		}
	}
	return HourlyData{}, errNotFound
}

type (
	ForecastOption func(*forecastOptions)

	forecastOptions struct{}
)

func Forecast(latitude, longitude float64, options ...ForecastOption) (*ForecastResponse, error) { // latitude=52.52&longitude=13.41
	var queryOpts forecastOptions

	for _, option := range options {
		option(&queryOpts)
	}

	hourlyOptions := []string{"temperature_2m", "relativehumidity_2m", "windspeed_10m"}

	res, err := resty.New().R().SetQueryParams(map[string]string{
		"hourly":          strings.Join(hourlyOptions, ","),
		"longitude":       fmt.Sprintf("%.f", longitude),
		"latitude":        fmt.Sprintf("%.f", latitude),
		"current_weather": "true",
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
