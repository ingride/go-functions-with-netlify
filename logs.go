package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"time"
)

func invocationLogRequest(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Println("a beautiful log line: ", time.Now())
	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       "Oh Hai Hello from your Logging Function!",
	}, nil
}

func main() {
	lambda.Start(invocationLogRequest)
}
