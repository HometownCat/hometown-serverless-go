package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"runtime"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"hometown.com/hometown-serverless-go/controller"
	"hometown.com/hometown-serverless-go/modules/common"
	"hometown.com/hometown-serverless-go/modules/database"
	"hometown.com/hometown-serverless-go/types"
)

func Handler(event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error){

	requestValid := common.RequestValid(&event,&[]types.ValidKey{
		{
			Key: "email",
			KeyType: "string",
		},
		{
			Key: "password",
			KeyType: "string",
		},
		{
			Key: "username",
			KeyType: "string",
		},
	})
	
	if !requestValid {
		response, err := common.ResponseError(errors.New("not avaliable request data [body]:[" + event.Body + "]"))
		return *response,err
	}
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

	defer database.MasterDatabase.Close()
	defer database.SlaveDatabase.Close()

	lambda.Start(Handler)
}