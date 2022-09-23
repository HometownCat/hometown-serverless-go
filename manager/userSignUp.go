package manager

import (
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
	token,tokenErr := handler.TokenGenerator(user)
	if tokenErr != nil {
		return nil, tokenErr
	}
	return token,nil
}