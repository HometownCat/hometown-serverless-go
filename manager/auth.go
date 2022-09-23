package manager

import (
	"encoding/json"

	"golang.org/x/crypto/bcrypt"
	"hometown.com/hometown-serverless-go/handler"
	"hometown.com/hometown-serverless-go/types"
)


func UserSignUp(user *types.User) (*types.TokenData, *error) {
	password,bcryptErr := bcrypt.GenerateFromPassword([]byte(user.Password),4)
	if bcryptErr != nil {
		return nil, &bcryptErr
	}
	user.Password = string(password)
	err := handler.SignUp(user)
	if err != nil {
		return nil, err
	}
	token,tokenErr := handler.TokenGenerator(&types.SendUserInfo{
		Email: user.Email,
		Name: user.Name,
	})
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