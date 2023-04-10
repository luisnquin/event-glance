package main

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/luisnquin/event-glance/internal/config"
	"github.com/luisnquin/event-glance/pkg/holiday"
	"github.com/luisnquin/event-glance/pkg/weather"
	"github.com/samber/lo"
)

func main() {
	// lambda.Start(action())

	config, err := config.New()
	if err != nil {
		panic(err)
	}

	ctx, session := context.Background(), session.New()

	forecastResponseChan := lo.Async(func() *weather.ForecastResponse {
		response, err := weather.Forecast(ctx, config.Weather.Latitude, config.Weather.Longitude)
		if err != nil {
			panic(err)
		}

		return response
	})

	holidaysChan := lo.Async(func() []holiday.Day {
		holidays, err := holiday.Search(ctx, time.Now().Year(), config.Holiday.CountryCode)
		if err != nil {
			panic(err)
		}

		return holidays
	})

	weatherInTwoDays, err := (<-forecastResponseChan).AfterCurrentWeather((time.Hour * 24) * 2)
	if err != nil {
		panic(err)
	}

	if !config.DebugMode {
		messageId, err := sendEmail(ctx, config, session, map[string]any{
			"weatherInTwoDays": weatherInTwoDays,
			"holidays":         <-holidaysChan,
		})
		if err != nil {
			panic(err)
		}

		fmt.Printf("{\"messageId\": %q}\n", messageId)
	}
}
