package manager

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"

	"hometown.com/hometown-serverless-go/handler"
	"hometown.com/hometown-serverless-go/types"
)


func UserSignUp(user *types.User) (*types.TokenData, *error) {
	hashPassword := sha256.Sum256([]byte(user.Password))
	user.Password = hex.EncodeToString(hashPassword[:])
	err := handler.SignUp(user)
	if err != nil {
		return nil, err
	}

	var sendUserInfo types.SendUserInfo
	userBin,jsonErr := json.Marshal(user)
	if jsonErr != nil {
		return nil, err
	}

	json.Unmarshal(userBin,&sendUserInfo)

	token,tokenErr := handler.TokenGenerator(&sendUserInfo)
	if tokenErr != nil {
		return nil, tokenErr
	}
	return token,nil
}

func TokenGenerator(email *string, password *string) (*map[string] interface{}, *error) {

	var userData map[string]interface{}

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