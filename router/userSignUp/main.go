package main

import (
	"encoding/json"
	"runtime"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"hometown.com/hometown-serverless-go/controller"
	"hometown.com/hometown-serverless-go/modules/common"
)

func Handler(event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error){
	token := controller.UserSignUp()

	bin,err := json.Marshal(&token)

	if err != nil {
		return common.ResponseError(err)
	}

	return events.APIGatewayProxyResponse{
		Body: string(bin),
		StatusCode: 200,
	}, nil
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	lambda.Start(Handler)
}