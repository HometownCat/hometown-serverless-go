package validation

import (
	"hometown.com/hometown-serverless-go/manager"
	"hometown.com/hometown-serverless-go/types"
)

func UserValidation(token *string, userInfo *types.SendUserInfo) error {
	if token == nil {
		return nil
	}	
	sendUserInfo := types.SendUserInfo{}
	parseErr := manager.AccessTokenParse(token,&sendUserInfo)

	if parseErr != nil {
		return parseErr
	}
	return nil
}
