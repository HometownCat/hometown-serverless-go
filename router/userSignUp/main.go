package main

import (
	"encoding/json"
	"runtime"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"hometown.com/hometown-serverless-go/controller"
	"hometown.com/hometown-serverless-go/modules/common"
	"hometown.com/hometown-serverless-go/types"
)

func Handler(event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error){
	token := controller.UserSignUp()
	responseData := types.ResponseData{
		Message: "success",
		Data: &token,
	}
	bin,err := json.Marshal(&responseData)

	if err != nil {
		response, err := common.ResponseError(err)
		return *response,*err
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