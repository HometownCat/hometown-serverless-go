package manager

import (
	"encoding/json"
	"os"

	"hometown.com/hometown-serverless-go/handler"
	"hometown.com/hometown-serverless-go/types"
)

func AccessTokenParse(token *string) (*types.SendUserInfo, error) {
	
	returnData := types.SendUserInfo{}
	secretKey := os.Getenv("JWT_ACCESS_SECRET_KEY")
	tokenData, parseErr := handler.TokenParser(token, &secretKey)
	if parseErr != nil {
		return  nil,parseErr
	}
	
	tokenBin,_ := json.Marshal(*tokenData)

	json.Unmarshal(tokenBin,&returnData)
	// returnData["id"] = tokenData.Id
	// returnData["email"] = tokenData.Email
	// returnData["username"] = tokenData.Username
	// returnData["address"] = tokenData.Address
	// returnData["phoneNumber"] = tokenData.PhoneNumber
	// returnData["profileImage"] = tokenData.ProfileImage

	return &returnData,nil
}