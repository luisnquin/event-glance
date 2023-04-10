package holiday

import (
	"encoding/json"
	"time"
)

// Represents a single holiday day, with information about the holiday name.
type Date struct {
	// Counties where the holiday is observed (if applicable).
	Counties []string `json:"counties"`
	// ISO-3166 country code for the country where the holiday is observed.
	CountryCode string `json:"countryCode"`
	// Date of the holiday.
	Date time.Time `json:"date"`
	// Indicates whether the holiday is fixed (i.e. always occurs on the same date).
	Fixed bool `json:"fixed"`
	// Indicates whether the holiday is observed globally.
	Global bool `json:"global"`
	// Year the holiday was first observed.
	LaunchYear uint16 `json:"launchYear"`
	// Local name of the holiday .
	LocalName string `json:"localName"`
	// Name of the holiday.
	Name string `json:"name"`
	// Holiday types.
	Types []string `json:"types"`
}

// A custom unmarshal method that helps to parse the holiday date.
func (d *Date) UnmarshalJSON(b []byte) error {
	type Alias Date

	aux := struct {
		Date string `json:"date"`
		*Alias
	}{
		Alias: (*Alias)(d),
	}

	err := json.Unmarshal(b, &aux)
	if err != nil {
		return err
	}

	t, err := time.Parse("2006-01-02", aux.Date)
	if err != nil {
		return err
	}

	d.CountryCode = aux.CountryCode
	d.LaunchYear = aux.LaunchYear
	d.LocalName = aux.LocalName
	d.Counties = aux.Counties
	d.Global = aux.Global
	d.Fixed = aux.Fixed
	d.Types = aux.Types
	d.Name = aux.Name
	d.Date = t

	return nil
}
