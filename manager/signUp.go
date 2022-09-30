package manager

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"

	"github.com/google/uuid"
	"hometown.com/hometown-serverless-go/handler"
	"hometown.com/hometown-serverless-go/types"
)

func UserSignUp(user *types.User) (*types.SendUserInfo, error) {
	hashPassword := sha256.Sum256([]byte(user.Password))
	user.Password = hex.EncodeToString(hashPassword[:])

	getUser, getUserErr := handler.GetUser(&user.Email, nil)

	if getUserErr != nil {
		return nil, getUserErr
	}

	if getUser != nil {
		return nil, errors.New("user aleady exist")
	}

	err := handler.SignUp(user)
	if err != nil {
		return nil, err
	}

	var sendUserInfo types.SendUserInfo

	userBin, jsonErr := json.Marshal(*user)

	if jsonErr != nil {
		return nil, err
	}

	json.Unmarshal(userBin, &sendUserInfo)

	accessToken, tokenErr := handler.RedisTokenGenerator(&sendUserInfo)
	revokeToken := uuid.NewString()

	if tokenErr != nil {
		return nil, tokenErr
	}

	redisErr := handler.SetUserToken(accessToken, &revokeToken, &sendUserInfo.Id)

	if redisErr != nil {
		return nil, redisErr
	}

	sendUserInfo.AccessToken = accessToken
	sendUserInfo.RevokeToken = &revokeToken

	return &sendUserInfo, nil
}
