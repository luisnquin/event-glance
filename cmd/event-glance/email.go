package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
	"github.com/luisnquin/event-glance/internal/config"
	"github.com/luisnquin/event-glance/template"
)

func sendEmail(ctx context.Context, config config.App, session *session.Session, data any) (string, error) {
	sesClient := ses.New(session,
		aws.NewConfig().WithRegion(config.Email.Region))

	tpl, err := template.Load()
	if err != nil {
		return "", err
	}

	otherBody, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	output, err := sesClient.SendEmailWithContext(ctx, &ses.SendEmailInput{
		Source: config.Email.Sender,
		Message: &ses.Message{
			Subject: &ses.Content{
				Charset: aws.String("UTF-8"),
				Data:    aws.String(tpl.Subject),
			},
			Body: &ses.Body{
				Html: &ses.Content{
					Charset: aws.String("UTF-8"),
					Data:    aws.String(fmt.Sprintf("%s %s", tpl.Body, otherBody)),
				},
			},
		},
		Destination: &ses.Destination{
			ToAddresses: config.Email.Receivers,
		},
	})
	if err != nil {
		return "", err
	}

	return *output.MessageId, nil
}
