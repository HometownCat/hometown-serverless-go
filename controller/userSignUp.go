package controller

import (
	"hometown.com/hometown-serverless-go/types"
)

func UserSignUp() *types.TokenData {
	// 테스트 데이터
	return &types.TokenData {
		AccessToken: "sdasdsaqfwefqw",
		RevokeToken: "fqwfwfqwqw",
	}
}