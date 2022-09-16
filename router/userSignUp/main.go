package main

import (
	"runtime"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func Handler(event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error){
	return events.APIGatewayProxyResponse{
		Body: "success",
		StatusCode: 200,
	}, nil
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	lambda.Start(Handler)
}