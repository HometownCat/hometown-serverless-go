package common

import (
	"github.com/aws/aws-lambda-go/events"
)

func IsError(err error) bool {
	return err != nil
}

func ResponseError(err error) (*events.APIGatewayProxyResponse, error){
	return &events.APIGatewayProxyResponse{
		Body: "{\"error\":\""+err.Error()+"\"}",
		StatusCode: 400,
	},nil
}