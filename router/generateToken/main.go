package main

import (
	"encoding/json"
	"runtime"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"hometown.com/hometown-serverless-go/controller"
	"hometown.com/hometown-serverless-go/modules/common"
	"hometown.com/hometown-serverless-go/modules/database"
	"hometown.com/hometown-serverless-go/types"
)

func Handler(event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error){

	userData,tokenErr := controller.GenerateToken(&event)

	if tokenErr != nil{
		errRes,err := common.ResponseError(tokenErr);
		return *errRes,err
	}

	responseData := types.ResponseData{
		Message: "success",
		Data: *userData,
	}
	bin,jsonErr := json.Marshal(responseData)

	if jsonErr != nil {
		errRes,err := common.ResponseError(jsonErr)
		return *errRes,err
	}

	return events.APIGatewayProxyResponse{
		Body: string(bin),
		StatusCode: 200,
	}, nil

}

func main(){
	runtime.GOMAXPROCS(runtime.NumCPU())

	defer database.MasterDatabase.Close()
	defer database.SlaveDatabase.Close()

	lambda.Start(Handler)
}