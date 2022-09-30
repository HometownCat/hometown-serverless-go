package manager

import (
	"encoding/json"

	"hometown.com/hometown-serverless-go/handler"
	"hometown.com/hometown-serverless-go/types"
)

func AccessTokenParse(token *string) (*types.SendUserInfo, error) {

	returnData := types.SendUserInfo{}
	tokenData, parseErr := handler.TokenParser(token)
	if parseErr != nil {
		return nil, parseErr
	}

	tokenBin, _ := json.Marshal(*tokenData)

	json.Unmarshal(tokenBin, &returnData)
	return &returnData, nil
}
