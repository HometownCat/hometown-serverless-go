package manager

import (
	"hometown.com/hometown-serverless-go/handler"
	"hometown.com/hometown-serverless-go/types"
)

func AccessTokenParse(token *string, userInfo *types.SendUserInfo) error {
	parseErr := handler.TokenParser(token,userInfo)
	if parseErr != nil {
		return parseErr
	}
	return nil
}
