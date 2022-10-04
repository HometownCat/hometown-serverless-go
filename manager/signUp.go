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

func UserSignUp(user *types.User, userInfo *types.SendUserInfo) error {
	hashPassword := sha256.Sum256([]byte(user.Password))
	user.Password = hex.EncodeToString(hashPassword[:])


	getUserErr := handler.GetUser(&user.Email, nil, userInfo)

	if getUserErr != nil {
		return getUserErr
	}
	if userInfo.Id != 0 {
		return errors.New("user aleady exist")
	}

	err := handler.SignUp(user)
	
	if err != nil {
		return err
	}

	userBin, _ := json.Marshal(*user)

	json.Unmarshal(userBin, &userInfo)
	tokenErr := handler.RedisTokenGenerator(userInfo)
	revokeToken := uuid.NewString()

	if tokenErr != nil {
		return tokenErr
	}

	redisErr := handler.SetUserToken(userInfo.AccessToken, &revokeToken, &userInfo.Id)

	if redisErr != nil {
		return redisErr
	}

	userInfo.RevokeToken = &revokeToken

	return nil
}
