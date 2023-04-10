package holiday_test

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/luisnquin/event-glance/pkg/holiday"
)

func TestSmokeForYear(t *testing.T) {
	ctx, currentYear := context.Background(), time.Now().Year()

	holidays, err := holiday.SearchForYear(ctx, "fr", currentYear)
	if err != nil {
		panic(err)
	}

	data, err := json.Marshal(holidays)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", data)
}

func TestSmokeForYears(t *testing.T) {
	ctx, currentYear := context.Background(), time.Now().Year()

	holidays, err := holiday.SearchForYears(ctx, "US", currentYear-6, currentYear, currentYear+5, currentYear+10)
	if err != nil {
		panic(err)
	}

	data, err := json.Marshal(holidays)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", data)
}

func TestSearch(t *testing.T) {
	ctx := context.Background()

	holidays, err := holiday.SearchForYear(ctx, "US", 2023)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	} else if len(holidays) == 0 {
		t.Errorf("expected at least one holiday, got none")
	}

	holidays, err = holiday.SearchForYear(ctx, "invalid", 2023)
	if err == nil {
		t.Error("expected error, got none")
		t.Fail()
	} else if len(holidays) > 0 {
		t.Errorf("expected zero holidays but got %d", len(holidays))
		t.Fail()
	}

	holidays, err = holiday.SearchForYear(ctx, "US", 0)
	if err == nil {
		t.Error("expected error, got none")
	} else if len(holidays) > 0 {
		t.Errorf("expected zero holidays but got %d", len(holidays))
		t.Fail()
	}
}
