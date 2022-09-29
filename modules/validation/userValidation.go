package validation

import (
	"hometown.com/hometown-serverless-go/manager"
	"hometown.com/hometown-serverless-go/types"
)

func UserValidation(token *string) (*types.SendUserInfo,error) {

	if token == nil{
		return nil, nil
	}
	userInfo, parseErr := manager.AccessTokenParse(token)

	if parseErr != nil {
		return nil, parseErr
	}
	return userInfo, nil
}