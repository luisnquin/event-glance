package config

import (
	"fmt"

	"github.com/caarlos0/env/v8"
	"github.com/luisnquin/event-glance/internal/validate"
)

type App struct {
	DebugMode bool `env:"DEBUG_MODE,required"`
	Holiday   HolidayServiceConfig
	Weather   WeatherServiceConfig
	Email     EmailServiceConfig
}

type HolidayServiceConfig struct {
	CountryCode string `env:"COUNTRY_CODE,required"`
}

type WeatherServiceConfig struct {
	Latitude  float64 `env:"LATITUDE,required"`
	Longitude float64 `env:"LONGITUDE,required"`
}

type EmailServiceConfig struct {
	Region    string    `env:"SES_REGION,required"`
	Sender    *string   `env:"SES_SENDER,required"`
	Receivers []*string `env:"SES_RECEIVERS,required" envSeparator:","`
}

func New() (App, error) {
	var app App

	if err := env.Parse(&app); err != nil {
		return App{}, err
	}

	if err := app.validate(); err != nil {
		return App{}, err
	}

	return app, nil
}

func (app App) validate() error {
	if !validate.LikeAWSRegion(app.Email.Region) {
		return fmt.Errorf("'%s' is not valid as an AWS region", app.Email.Region)
	}

	if !validate.IsEmail(*app.Email.Sender) {
		return fmt.Errorf("the sender email '%s' isn't valid", *app.Email.Sender)
	}

	for _, receiver := range app.Email.Receivers {
		if !validate.IsEmail(*receiver) {
			return fmt.Errorf("the receiver email '%s' isn't valid", *app.Email.Sender)
		}
	}

	return nil
}
