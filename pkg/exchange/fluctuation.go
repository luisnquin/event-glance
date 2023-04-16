package exchange

import (
	"context"
	"encoding/json"
	"errors"
	"net/url"
	"strings"
	"time"

	"github.com/luisnquin/event-glance/pkg/exchange/currency"
)

type (
	FluctuationResponse struct {
		// Indicates whether there was any fluctuation in the exchange rates during the specified period
		Fluctuation bool `json:"fluctuation"`
		// The base currency
		Base string `json:"base"`
		// the start date of the period for which the rates are being provided
		StartDate time.Time `json:"start_date"`
		// The end date of the period for which the rates are being provided
		EndDate time.Time `json:"end_date"`
		// The exchange rate information for the specified currencies
		Rates map[string]FluctuationRate `json:"rates"`
	}

	FluctuationRate struct {
		// The absolute change in the exchange rate
		Change float64 `json:"change"`
		// Specifies the percentage change in the exchange rate
		ChangePercentage float64 `json:"change_pct"`
		// Specifies the exchange rate of 'symbols' against the base currency at the beginning of the specified period
		StartRate float64 `json:"start_rate"`
		// Specifies the exchange rate of 'symbols' against the base currency at the end of the specified period
		EndRate float64 `json:"end_rate"`
	}
)

func (f *FluctuationResponse) UnmarshalJSON(data []byte) error {
	type Alias FluctuationResponse

	aux := struct {
		StartDate string `json:"start_date"`
		EndDate   string `json:"end_date"`
		*Alias
	}{
		Alias: (*Alias)(f),
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	startDate, err := time.Parse(time.DateOnly, aux.StartDate)
	if err != nil {
		return err
	}

	endDate, err := time.Parse(time.DateOnly, aux.EndDate)
	if err != nil {
		return err
	}

	f.Fluctuation = aux.Fluctuation
	f.StartDate = startDate
	f.EndDate = endDate
	f.Rates = aux.Rates
	f.Base = aux.Base

	return nil
}

// Sends a request to the exchange rate API to calculate exchange rate fluctuations
// between a base currency and one or more specified currencies over a given time period.
func Fluctuation(ctx context.Context, apiKey string, startDate, endDate time.Time, base string, toCompare []string) (FluctuationResponse, error) {
	if len(toCompare) == 0 {
		return FluctuationResponse{}, errors.New("'toCompare' must have at least one iso4217 currency code")
	}

	for _, code := range toCompare {
		if !currency.Is(code) {
			return FluctuationResponse{}, notValidCurrencyCode(code)
		}
	}

	if !currency.Is(base) {
		return FluctuationResponse{}, notValidCurrencyCode(base)
	}

	query := url.Values{
		"start_date": []string{startDate.Format(time.DateOnly)},
		"end_date":   []string{endDate.Format(time.DateOnly)},
		"base":       []string{base},
		"symbols":    []string{strings.Join(toCompare, ",")},
	}

	return doGetRequest[FluctuationResponse](ctx, query, "/fluctuation", apiKey)
}
