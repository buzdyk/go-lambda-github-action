package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"testing"
)

func TestHandler(t *testing.T) {
	t.Run("Successful response", func(t *testing.T) {
		response, err := handler(events.APIGatewayProxyRequest{
			RequestContext: events.APIGatewayProxyRequestContext{
				Identity: events.APIGatewayRequestIdentity{SourceIP: "127.0.0.1"},
			},
		})

		if err != nil {
			t.Fatal("Everything should be ok")
		}

		expected := "Hello, 127.0.0.1!"

		if response.Body != expected {
			t.Fatal(fmt.Printf("Response message is wrong. Expected %s, got %s", expected, response.Body))
		}
	})
}
