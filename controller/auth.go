package controller

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"hometown.com/hometown-serverless-go/manager"
	"hometown.com/hometown-serverless-go/types"
)

func UserSignUp(event *events.APIGatewayProxyRequest) (*map[string] interface{}, *error) {
	var user types.User
	jsonErr := json.Unmarshal([]byte(event.Body),&user)
	if jsonErr != nil {
		return nil, &jsonErr
	}

	token,err := manager.UserSignUp(&user)

	if err != nil {
		return nil,err
	}

	var userData map[string] interface{}

	tokenBin, _ := json.Marshal(&token)
	json.Unmarshal(tokenBin,&userData)

	userData["email"] = user.Email
	userData["name"] = user.Name


	// 테스트 데이터
	return &userData, nil
}

func GenerateToken(event *events.APIGatewayProxyRequest) (*map[string] interface{}, *error) {
	var reqUser types.LoginUser
	jsonErr := json.Unmarshal([]byte(event.Body),&reqUser)

	if jsonErr != nil {
		return nil, &jsonErr
	}
	userData,err := manager.TokenGenerator(&reqUser.Email, &reqUser.Password)
	if err != nil {
		return nil, err
	}
	return userData,nil
}