package controller

import (
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"hometown.com/hometown-serverless-go/manager"
	"hometown.com/hometown-serverless-go/types"
)

func UserSignUp(event *events.APIGatewayProxyRequest) (*types.SendUserInfo, error) {
	var user types.User
	jsonErr := json.Unmarshal([]byte(event.Body),&user)
	user.UserIp = event.RequestContext.Identity.SourceIP
	if jsonErr != nil {
		return nil, jsonErr
	}
	sendUserInfo,err := manager.UserSignUp(&user)

	if err != nil {
		fmt.Println("here")
		return nil,err
	}
	return sendUserInfo, nil
}
