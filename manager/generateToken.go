package manager

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"

	"hometown.com/hometown-serverless-go/handler"
	"hometown.com/hometown-serverless-go/types"
)

func TokenGenerator(email *string, password *string) (*types.SendUserInfo, error) {
	hashPassword := sha256.Sum256([]byte(*password))
	*password = hex.EncodeToString(hashPassword[:])
	var userData types.SendUserInfo

	userInfo, getUserErr := handler.GetUser(email, password)
	if getUserErr != nil {
		return nil, getUserErr
	}
	token, tokenErr := handler.TokenGenerator(userInfo)

	if tokenErr != nil {
		return nil, tokenErr
	}

	userBin, _ := json.Marshal(&userInfo)
	tokenBin, _ := json.Marshal(&token)

	json.Unmarshal(userBin,&userData)
	json.Unmarshal(tokenBin,&userData)
	
	return &userData,nil
}