package main

import (
	"encoding/json"
	"runtime"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"hometown.com/hometown-serverless-go/controller"
	"hometown.com/hometown-serverless-go/modules/common"
	"hometown.com/hometown-serverless-go/modules/database"
	"hometown.com/hometown-serverless-go/modules/redis"
	"hometown.com/hometown-serverless-go/types"
)

func Handler(event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	params, requestErr := common.RequestValid(&event, []types.ValidKey{
		{
			Key:     "email",
			KeyType: "string",
		},
		{
			Key:     "password",
			KeyType: "string",
		},
		{
			Key:     "username",
			KeyType: "string",
		},
	})

	var response events.APIGatewayProxyResponse

	if requestErr != nil {
		err := common.ResponseError(requestErr, &response)
		return response, err
	}

	var userInfo types.SendUserInfo

	err := controller.UserSignUp(&params, &userInfo)

	if err != nil {
		err := common.ResponseError(err, &response)
		return response, err
	}

	responseData := types.ResponseData{
		Message: "success",
		Data:    userInfo,
	}

	bin, jsonErr := json.Marshal(&responseData)

	if jsonErr != nil {
		err := common.ResponseError(jsonErr, &response)
		return response, err
	}

	response.Body = string(bin)
	response.StatusCode = 200

	return response, nil
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	defer database.MasterDatabase.Close()
	defer database.SlaveDatabase.Close()
	defer redis.Close()
	
	lambda.Start(Handler)
}
