package main

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
)

type LambdaHandlerFunc func(ctx context.Context, request events.APIGatewayV2HTTPRequest) (*events.APIGatewayProxyResponse, error)

func action() LambdaHandlerFunc {
	return func(ctx context.Context, request events.APIGatewayV2HTTPRequest) (*events.APIGatewayProxyResponse, error) {
		return nil, nil
	}
}
