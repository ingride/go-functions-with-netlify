package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"time"
)

func invocationLogRequest(request events.APIGatewayProxyRequest)  {
	fmt.Println("a beautiful log line: ", time.Now())
}

func main() {
	lambda.Start(invocationLogRequest)
}
