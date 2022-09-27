package handler

import (
	"errors"
	"sync"

	"github.com/google/uuid"
	"hometown.com/hometown-serverless-go/types"
)


func SignUp(user *types.User) *error {
	// db 처리 추가
	return nil
}

func GetUser(email *string , password *string) (*types.SendUserInfo, *error){
	// db 처리 추가
	if false {
		passwordFailErr := errors.New("password not matched")
		return nil, &passwordFailErr
	} 
	return &types.SendUserInfo{
		Email: "test@naver.com",
		Name: "park",
	}, nil
} 

func TokenGenerator(user *types.SendUserInfo) (*types.TokenData, *error){
	// 토큰 발급 추가
	var accessToken string
	var revokeToken string

	var wg sync.WaitGroup

	wg.Add(2)

	go func(){
		defer wg.Done()
		accessToken = uuid.NewString()
	}()
	go func(){
		defer wg.Done()
		revokeToken = uuid.NewString()
	}()
	wg.Wait()
	return &types.TokenData{
		AccessToken: accessToken,
		RevokeToken: revokeToken,
	}, nil
}