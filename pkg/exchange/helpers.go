package exchange

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/goccy/go-json"
)

func doGetRequest[T any](ctx context.Context, query url.Values, path, apiKey string) (T, error) {
	var response T

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://api.apilayer.com/exchangerates_data"+path, http.NoBody)
	if err != nil {
		return response, err
	}

	req.URL.RawQuery = query.Encode()
	req.Header.Add(API_KEY, apiKey)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return response, err
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		body, err := io.ReadAll(res.Body)
		if err != nil {
			return response, err
		}

		return response, fmt.Errorf("unexpected status code: %d, body: %s", res.StatusCode, body)
	}

	return response, json.NewDecoder(res.Body).Decode(&response)
}
