package exchange

import (
	"context"
	"fmt"
	"net/url"
	"time"

	"github.com/goccy/go-json"
)

type (
	ConvertResponse struct {
		Date       time.Time            `json:"date"`
		Historical bool                 `json:"historical"`
		Info       ConvertResponseInfo  `json:"info"`
		Query      ConvertResponseQuery `json:"query"`
		Result     float64              `json:"result"`
	}

	ConvertResponseInfo struct {
		Rate      float64 `json:"rate"`
		TimeStamp float64 `json:"time_stamp"`
	}

	ConvertResponseQuery struct {
		Amount float64 `json:"amount"`
		From   string  `json:"from"`
		To     string  `json:"to"`
	}
)

func (c *ConvertResponse) UnmarshalJSON(data []byte) error {
	type Alias ConvertResponse

	aux := struct {
		Date string `json:"date"`
		*Alias
	}{
		Alias: (*Alias)(c),
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	t, err := time.Parse(time.DateOnly, aux.Date)
	if err != nil {
		return err
	}

	c.Date = t
	c.Historical = aux.Historical
	c.Result = aux.Result
	c.Query = aux.Query
	c.Info = aux.Info

	return nil
}

func Convert(ctx context.Context, apiKey string, date time.Time, amount float64, from, to string) (ConvertResponse, error) {
	query := url.Values{
		"from":   []string{from},
		"to":     []string{to},
		"amount": []string{fmt.Sprintf("%.f", amount)},
		"date":   []string{date.Format(time.DateOnly)},
	}

	return doGetRequest[ConvertResponse](ctx, query, "/convert", apiKey)
}
