// Library that provides the ability to search for public holidays in a specific
// country for a given year using an external REST API(https://github.com/nager/Nager.Date).
package holiday

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/goccy/go-json"
)

// Searches the holidays of an specific year and country using an external API
// and returns a slice of Day.
func SearchForYear(ctx context.Context, countryCode string, year int) ([]Date, error) {
	countryCode = strings.ToUpper(strings.TrimSpace(countryCode))
	if len(countryCode) != 2 {
		return nil, errors.New("invalid country code")
	}

	return searchForYear(ctx, countryCode, year)
}

// Allows to search for public holidays in a specific country for multiple years
// at once using an external REST API. All requests are made by goroutines.
func SearchForYears(ctx context.Context, countryCode string, years ...int) ([]Date, error) {
	countryCode = strings.ToUpper(strings.TrimSpace(countryCode))
	if len(countryCode) != 2 {
		return nil, errors.New("invalid country code")
	}

	if len(years) == 0 {
		return nil, errors.New("at least one year is required")
	}

	yearsSet := make(map[int]struct{}, len(years))

	for _, year := range years {
		yearsSet[year] = struct{}{}
	}

	holidaysChan := make(chan []Date)
	errsChan := make(chan error)

	for year := range yearsSet {
		go func(year int) {
			holidays, err := searchForYear(ctx, countryCode, year)
			if err != nil {
				errsChan <- err
			}

			holidaysChan <- holidays
		}(year)
	}

	var holidays []Date

	for i := 0; i < len(yearsSet); i++ {
		select {
		case holiday := <-holidaysChan:
			holidays = append(holidays, holiday...)
		case err := <-errsChan:
			if err != nil {
				return nil, err
			}
		}
	}

	return holidays, nil
}

func searchForYear(ctx context.Context, countryCode string, year int) ([]Date, error) {
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

	var holidays []Date

	return holidays, json.Unmarshal(resBody, &holidays)
}
