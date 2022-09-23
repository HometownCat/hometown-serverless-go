package common

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
)

func IsError(err error) bool {
	return err != nil
}

func ResponseError(err error) (*events.APIGatewayProxyResponse, *error){
	errBin, err := json.Marshal(err)

	if err != nil {
		return &events.APIGatewayProxyResponse{
			Body: "json parsing error",
			StatusCode: 400,
		},&err
	}

	return &events.APIGatewayProxyResponse{
		Body: string(errBin),
		StatusCode: 400,
	},&err
}