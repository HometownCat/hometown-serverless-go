package manager

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"sync"

	"hometown.com/hometown-serverless-go/handler"
	"hometown.com/hometown-serverless-go/types"
)


func UserSignUp(user *types.User) (*types.SendUserInfo, error) {
	hashPassword := sha256.Sum256([]byte(user.Password))
	user.Password = hex.EncodeToString(hashPassword[:])

	getUser, getUserErr := handler.GetUser(&user.Email,nil)

	if getUserErr != nil {
		return nil,getUserErr
	}


	if getUser != nil {
		return nil, errors.New("user aleady exist")
	}

	err := handler.SignUp(user)
	if err != nil {
		return nil, err
	}

	var sendUserInfo types.SendUserInfo

	userBin,jsonErr := json.Marshal(*user)
	
	if jsonErr != nil {
		return nil, err
	}

	json.Unmarshal(userBin,&sendUserInfo)

	var tokenErr error
	var wg sync.WaitGroup

	wg.Add(2)

	go func(){
		accessToken, accessErr := handler.TokenGenerator(&sendUserInfo,&JWT_ACCESS_SECRET_KEY,&JWT_ACCESS_AVALIABLE_TIME)
		sendUserInfo.AccessToken = accessToken
		if accessErr != nil {
			tokenErr = accessErr
		}
		defer wg.Done()
	}()
	
	go func(){
		revokeToken, revokeErr := handler.TokenGenerator(&sendUserInfo,&JWT_ACCESS_SECRET_KEY,&JWT_ACCESS_AVALIABLE_TIME)
		sendUserInfo.RevokeToken = revokeToken
		if revokeErr != nil {
			tokenErr = revokeErr
		}
		defer wg.Done()
	}()

	wg.Wait()
	
	if tokenErr != nil {
		return nil, tokenErr
	}

	return &sendUserInfo,nil
}
