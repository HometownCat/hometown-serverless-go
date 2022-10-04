package controller

import (
	"encoding/json"

	"hometown.com/hometown-serverless-go/manager"
	"hometown.com/hometown-serverless-go/types"
)

func UserSignUp(params *map[string]interface{}, userInfo *types.SendUserInfo) error {
	var user types.User

	paramsBin, _ := json.Marshal(params)
	jsonErr := json.Unmarshal(paramsBin, &user)
	if jsonErr != nil {
		return jsonErr
	}
	err := manager.UserSignUp(&user, userInfo)

	if err != nil {
		return err
	}
	return nil
}
