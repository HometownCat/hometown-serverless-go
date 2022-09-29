package manager

import (
	"os"

	"hometown.com/hometown-serverless-go/handler"
)

func AccessTokenParse(token *string) (*map[string]interface{}, error) {
	
	var returnData map[string]interface{}
	secretKey := os.Getenv("JWT_ACCESS_SECRET_KEY")
	tokenData, parseErr := handler.TokenParser(token, &secretKey)

	if parseErr != nil {
		return  nil,parseErr
	}
	
	returnData["id"] = tokenData.Id
	returnData["email"] = tokenData.Email
	returnData["username"] = tokenData.Username
	returnData["address"] = tokenData.Address
	returnData["phoneNumber"] = tokenData.PhoneNumber
	returnData["profileImage"] = tokenData.ProfileImage

	return &returnData,nil
}