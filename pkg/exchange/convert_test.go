package exchange_test

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/luisnquin/event-glance/pkg/exchange"
	"github.com/luisnquin/event-glance/pkg/exchange/currency"
)

func TestConvertSmoke(t *testing.T) {
	ctx, apiKey := context.Background(), os.Getenv("API_LAYER_API_KEY")
	now := time.Now()

	if apiKey == "" {
		t.Log("'API_LAYER_API_KEY' not provided, test skipped")

		return
	}

	response, err := exchange.Convert(ctx, apiKey, now, 1, currency.EUR, currency.USD)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	t.Log(response)
}
