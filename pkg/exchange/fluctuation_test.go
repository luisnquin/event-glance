package exchange_test

import (
	"context"
	"encoding/json"
	"os"
	"testing"
	"time"

	"github.com/luisnquin/event-glance/pkg/exchange"
	"github.com/luisnquin/event-glance/pkg/exchange/currency"
)

func TestFluctuationSmoke(t *testing.T) {
	ctx, apiKey := context.Background(), os.Getenv("API_LAYER_API_KEY")
	now := time.Now()

	if apiKey == "" {
		t.Log("'API_LAYER_API_KEY' not provided, test skipped")

		return
	}

	const day = time.Hour * 24

	response, err := exchange.Fluctuation(ctx, apiKey, now.Add(-day*30), now, currency.EUR, []string{currency.PEN})
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	data, _ := json.MarshalIndent(response, "", "\t")

	t.Logf("%s", data)
}
