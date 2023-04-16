package exchange_test

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/luisnquin/event-glance/pkg/exchange"
)

func TestFluctuationSmoke(t *testing.T) {
	ctx, apiKey := context.Background(), os.Getenv("API_LAYER_API_KEY")
	now := time.Now()

	response, err := exchange.Fluctuation(ctx, apiKey, now, now, "USD", []string{"EUR"})
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	t.Log(response)
}
