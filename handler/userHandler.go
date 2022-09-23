package handler

import (
	"github.com/google/uuid"
	"hometown.com/hometown-serverless-go/types"
)


func SignUp(user *types.User) *error {
	// db 처리 추가
	return nil
}

func TokenGenerator(user *types.User) (*types.TokenData, *error){
	// 토큰 발급 추가
	accessToken := uuid.NewString()
	revokeToken := uuid.NewString()
	return &types.TokenData{
		AccessToken: accessToken,
		RevokeToken: revokeToken,
	}, nil
}