package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type Event struct {
	Name string `json:"name"`
}

func handler(event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	message := fmt.Sprintf("Hello, %s!", event.RequestContext.Identity.SourceIP)
	return events.APIGatewayProxyResponse{
		Body:       message,
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(handler)
}
