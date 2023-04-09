package weather

import (
	"time"

	"github.com/goccy/go-json"
)

type (
	CurrentWeather struct {
		Temperature   float64   `json:"temperature"`
		WindSpeed     float64   `json:"windspeed"`
		WindDirection float64   `json:"winddirection"`
		WeatherCode   float64   `json:"weather_code"`
		IsDay         uint8     `json:"is_day"`
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
	}
)

func (h *CurrentWeather) UnmarshalJSON(b []byte) error {
	type Alias CurrentWeather

	aux := struct {
		Time string `json:"time"`
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

	return nil
}

func (h *Hourly) UnmarshalJSON(b []byte) error {
	type Alias Hourly

	aux := struct {
		Time []string `json:"time"`
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

	return nil
}
