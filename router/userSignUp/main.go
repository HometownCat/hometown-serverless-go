package main

import (
	"encoding/json"
	"fmt"
	"runtime"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"hometown.com/hometown-serverless-go/controller"
	"hometown.com/hometown-serverless-go/modules/common"
	"hometown.com/hometown-serverless-go/types"
)

func Handler(event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error){
	fmt.Println(event.RequestContext.Identity)
	userData, err := controller.UserSignUp(&event)
	if err != nil {
		response, err := common.ResponseError(err)
		return *response,err
	}
	responseData := types.ResponseData{
		Message: "success",
		Data: *userData,
	}
	bin,jsonErr := json.Marshal(&responseData)
	fmt.Println(responseData)
	if jsonErr != nil {
		response, err := common.ResponseError(jsonErr)
		return *response,err
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