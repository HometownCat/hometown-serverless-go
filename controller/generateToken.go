package controller

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"hometown.com/hometown-serverless-go/manager"
	"hometown.com/hometown-serverless-go/types"
)

func GenerateToken(event *events.APIGatewayProxyRequest) (*types.SendUserInfo, error) {
	var reqUser types.LoginUser
	jsonErr := json.Unmarshal([]byte(event.Body),&reqUser)

	if jsonErr != nil {
		return nil, jsonErr
	}
	userData,err := manager.TokenGenerator(&reqUser.Email, &reqUser.Password)
	if err != nil {
		return nil, err
	}
	return userData,nil
}