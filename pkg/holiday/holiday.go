package holiday

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/goccy/go-json"
)

// curl https://date.nager.at/api/v3/publicholidays/2023/PE

type Day struct {
	Date        time.Time `json:"date"`
	LocalName   string    `json:"localName"`
	Name        string    `json:"name"`
	CountryCode string    `json:"countryCode"`
	Fixed       bool      `json:"fixed"`
	Global      bool      `json:"global"`
	Counties    any       `json:"counties"`
	LaunchYear  uint16    `json:"launchYear"`
	Types       []string  `json:"types"`
}

func (d *Day) UnmarshalJSON(b []byte) error {
	type Alias Day

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

func Search(ctx context.Context, year int, countryCode string) ([]Day, error) {
	countryCode = strings.ToUpper(strings.TrimSpace(countryCode))

	url := fmt.Sprintf("https://date.nager.at/api/v3/publicholidays/%d/%s", year, countryCode)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, http.NoBody)
	if err != nil {
		return nil, err
	}

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

	var holidays []Day

	return holidays, json.Unmarshal(resBody, &holidays)
}
