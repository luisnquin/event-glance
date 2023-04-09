package config

import (
	"fmt"

	"github.com/caarlos0/env/v8"
	"github.com/luisnquin/event-glance/internal/validate"
)

type App struct {
	DebugMode    bool `env:"DEBUG_MODE,required"`
	EmailService EmailServiceConfig
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
	if !validate.LikeAWSRegion(app.EmailService.Region) {
		return fmt.Errorf("'%s' is not valid as an AWS region", app.EmailService.Region)
	}

	if !validate.IsEmail(*app.EmailService.Sender) {
		return fmt.Errorf("the sender email '%s' isn't valid", *app.EmailService.Sender)
	}

	for _, receiver := range app.EmailService.Receivers {
		if !validate.IsEmail(*receiver) {
			return fmt.Errorf("the receiver email '%s' isn't valid", *app.EmailService.Sender)
		}
	}

	return nil
}
