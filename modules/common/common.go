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

func ReturnNotNil(arg1 interface{}, arg2 interface{}) *interface{} {
	if arg1 != nil {
		return &arg1
	} else if arg2 != nil {
		return &arg2
	} else {
		return nil
	}
}