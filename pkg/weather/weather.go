package weather

import (
	"errors"
	"time"

	"github.com/goccy/go-json"
)

type (
	CurrentWeather struct {
		Temperature   float64   `json:"temperature"`
		WindSpeed     float64   `json:"windspeed"`
		WindDirection float64   `json:"winddirection"`
		WeatherCode   Code      `json:"weathercode"`
		IsDay         bool      `json:"is_day"`
		Time          time.Time `json:"time"`
	}

	HourlyUnits struct {
		Time               string `json:"time"`
		Temperature2M      string `json:"temperature_2m"`
		RelativeHumidity2M string `json:"relativehumidity_2m"`
		Windspeed10M       string `json:"windspeed_10m"`
	}

	Hourly struct {
		Time               []time.Time `json:"time"`
		Temperature2M      []float64   `json:"temperature_2m"`
		RelativeHumidity2M []float64   `json:"relativehumidity_2m"`
		WindSpeed10M       []float64   `json:"windspeed_10m"`
		WeatherCode        []Code      `json:"weathercode"`
		Cloudcover         []float64   `json:"cloudcover"`
		Visibility         []float64   `json:"visibility"`
		IsDay              []bool      `json:"is_day"`
	}
)

type HourlyData struct {
	Time               time.Time `json:"time"`
	Temperature2M      float64   `json:"temperature_2m"`
	RelativeHumidity2M float64   `json:"relativehumidity_2m"`
	WindSpeed10M       float64   `json:"windspeed_10m"`
	WeatherCode        Code      `json:"weathercode"`
	Cloudcover         float64   `json:"cloudcover"`
	Visibility         float64   `json:"visibility"`
	IsDay              bool      `json:"is_day"`
}

var errNotFound = errors.New("weather not found")

func IsNotFound(err error) bool {
	return errors.Is(err, errNotFound)
}

func (h *CurrentWeather) UnmarshalJSON(b []byte) error {
	type Alias CurrentWeather

	aux := struct {
		Time  string `json:"time"`
		IsDay uint8  `json:"is_day"`
		*Alias
	}{
		Alias: (*Alias)(h),
	}

	err := json.Unmarshal(b, &aux)
	if err != nil {
		return err
	}

	h.Time, err = time.Parse("2006-01-02T15:04", aux.Time)
	if err != nil {
		return err
	}

	if aux.IsDay == 1 {
		h.IsDay = true
	}

	h.WindDirection = aux.WindDirection
	h.WeatherCode = aux.WeatherCode
	h.Temperature = aux.Temperature
	h.WindSpeed = aux.WindSpeed

	return nil
}

func (h *Hourly) UnmarshalJSON(b []byte) error {
	type Alias Hourly

	aux := struct {
		Time  []string `json:"time"`
		IsDay []uint8  `json:"is_day"`
		*Alias
	}{}

	err := json.Unmarshal(b, &aux)
	if err != nil {
		return err
	}

	h.Time = make([]time.Time, len(aux.Time))

	for i, rawTime := range aux.Time {
		t, err := time.Parse("2006-01-02T15:04", rawTime)
		if err != nil {
			return err
		}

		h.Time[i] = t
	}

	h.IsDay = make([]bool, len(aux.IsDay))

	for i, isDay := range aux.IsDay {
		if isDay == 1 {
			h.IsDay[i] = true
		} else {
			h.IsDay[i] = false
		}
	}

	h.RelativeHumidity2M = aux.RelativeHumidity2M
	h.Temperature2M = aux.Temperature2M
	h.WindSpeed10M = aux.WindSpeed10M
	h.WeatherCode = aux.WeatherCode
	h.Cloudcover = aux.Cloudcover
	h.Visibility = aux.Visibility

	return nil
}
