package holiday_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/goccy/go-json"
	"github.com/luisnquin/event-glance/pkg/holiday"
)

func TestSmoke(t *testing.T) {
	ctx := context.Background()

	currentYear := time.Now().Year()

	holidays, err := holiday.Search(ctx, currentYear, "fr")
	if err != nil {
		panic(err)
	}

	data, err := json.Marshal(holidays)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", data)
}
