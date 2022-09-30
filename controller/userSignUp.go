package controller

import (
	"encoding/json"
	"fmt"

	"hometown.com/hometown-serverless-go/manager"
	"hometown.com/hometown-serverless-go/types"
)

func UserSignUp(params *map[string]interface{}) (*types.SendUserInfo, error) {
	var user types.User

	paramsBin,_ := json.Marshal(params)
	jsonErr := json.Unmarshal(paramsBin,&user)
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
