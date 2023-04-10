package weather

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/goccy/go-json"
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
				WeatherCode:        f.Hourly.WeatherCode[index],
				Cloudcover:         f.Hourly.Cloudcover[index],
				Visibility:         f.Hourly.Visibility[index],
				IsDay:              f.Hourly.IsDay[index],
				Time:               f.Hourly.Time[index],
			}, nil
		}
	}

	return HourlyData{}, errNotFound
}

type (
	ForecastOption func(*forecastOptions)

	forecastOptions struct {
		precipitationUnit string
		temperatureUnit   string
		windspeedUnit     string
		elevation         *float64
	}
)

func Temperature(unit TemperatureUnit) ForecastOption {
	return func(fo *forecastOptions) {
		fo.temperatureUnit = string(unit)
	}
}

func Precipitation(unit PrecipitationUnit) ForecastOption {
	return func(fo *forecastOptions) {
		fo.precipitationUnit = string(unit)
	}
}

func Windspeed(unit WindspeedUnit) ForecastOption {
	return func(fo *forecastOptions) {
		fo.windspeedUnit = string(unit)
	}
}

func Elevation(elevation float64) ForecastOption {
	return func(fo *forecastOptions) {
		fo.elevation = &elevation
	}
}

func Forecast(ctx context.Context, latitude, longitude float64, options ...ForecastOption) (*ForecastResponse, error) {
	var queryOpts forecastOptions

	for _, option := range options {
		option(&queryOpts)
	}

	query := make(url.Values)
	query.Add("hourly", "temperature_2m,relativehumidity_2m,windspeed_10m,cloudcover,is_day,visibility,weathercode")
	query.Add("longitude", fmt.Sprintf("%.f", longitude))
	query.Add("latitude", fmt.Sprintf("%.f", latitude))
	query.Add("current_weather", "true")

	if queryOpts.elevation != nil {
		query.Add("elevation", fmt.Sprintf("%.f", *queryOpts.elevation))
	}

	if queryOpts.temperatureUnit != "" {
		query.Add("temperature_unit", queryOpts.temperatureUnit)
	}

	if queryOpts.windspeedUnit != "" {
		query.Add("windspeed_unit", queryOpts.windspeedUnit)
	}

	if queryOpts.precipitationUnit != "" {
		query.Add("precipitation_unit", queryOpts.precipitationUnit)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://api.open-meteo.com/v1/forecast", http.NoBody)
	if err != nil {
		return nil, err
	}

	req.URL.RawQuery = query.Encode()

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if err := res.Body.Close(); err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d, body: %s", res.StatusCode, resBody)
	}

	var forecast ForecastResponse

	return &forecast, json.Unmarshal(resBody, &forecast)
}
