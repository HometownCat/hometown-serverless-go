package controller

import (
	"encoding/json"

	"hometown.com/hometown-serverless-go/manager"
	"hometown.com/hometown-serverless-go/types"
)

func GenerateToken(params *map[string]interface{}, userInfo *types.SendUserInfo) error {
	var reqUser types.LoginUser

	paramsBin, _ := json.Marshal(*params)
	jsonErr := json.Unmarshal(paramsBin, &reqUser)

	if jsonErr != nil {
		return jsonErr
	}

	err := manager.TokenGenerator(&reqUser.Email, &reqUser.Password, userInfo)

	if err != nil {
		return err
	}

	return nil
}

// func Gentoken(event *events.APIGatewayProxyRequest, callback func(error, *types.SendUserInfo)) {
// 	var reqUser types.LoginUser
// 	json.Unmarshal([]byte(event.Body),&reqUser)
// 	userData,err := manager.TokenGenerator(&reqUser.Email, &reqUser.Password)
// 	callback(err,userData)
// }
