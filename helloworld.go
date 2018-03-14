package main

import (
	"time"
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/thundra-io/thundra-lambda-agent-go/thundra"
	"github.com/thundra-io/thundra-lambda-agent-go/trace"
)

type MyEvent struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type MyResponse struct {
	Message string `json:"answer:"`
}

// This is your lambda function
func HandleLambdaEvent(event MyEvent) (MyResponse, error) {
	time.Sleep(15 * time.Millisecond)
	return MyResponse{Message: fmt.Sprintf("%s is %d years old!", event.Name, event.Age)}, nil
}

func main() {
	// Instantiate Thundra Agent with Trace Support
	t := thundra.NewBuilder().
			AddPlugin(&trace.Trace{}).
			SetAPIKey(/*TODO login https://console.thundra.io to get your APIKey*/).
			Build()

	// Wrap your lambda function with Thundra
	lambda.Start(thundra.Wrap(HandleLambdaEvent, t))
}