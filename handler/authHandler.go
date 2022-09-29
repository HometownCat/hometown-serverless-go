package handler

import (
	"errors"
	"sync"

	"github.com/google/uuid"
	"hometown.com/hometown-serverless-go/modules/database"
	"hometown.com/hometown-serverless-go/types"
)


func SignUp(user *types.User) error {
	// db 처리 추가
	strQuery := "INSERT INTO `user` SET username = \"" + user.Username +  "\", email = \"" + user.Email + "\", password = \"" + user.Password + "\", userIp = \"" + user.UserIp +"\", status = 0"
	if user.Address != nil {
		strQuery += ", address = " + *user.Address
	}
	if user.PhoneNumber != nil {
		strQuery += ", phoneNumber = " + *user.PhoneNumber
	}
	if user.ProfileImage != nil {
		strQuery += ", profileImage = " + *user.ProfileImage
	}
	result,conErr := database.MasterDatabase.Exec(strQuery)

	if conErr == nil {
		id,rowErr := result.LastInsertId()
		conErr = rowErr
		user.Id = &id
	}
	defer database.MasterDatabase.Close()
	return conErr
}

func GetUser(email *string , password *string) (*types.SendUserInfo, error){
	// db 처리 추가

	var sendUserInfo []types.SendUserInfo
	var passwordFailErr error
	strQuery := "SELECT id, email, username, address, phoneNumber, profileImage FROM `user` WHERE `email` = \"" + *email + "\""

	if password != nil {
		passwordFailErr = errors.New("password not matched");
		strQuery += ", `password` = \"" + *password + "\""
	}
	conErr := database.MasterDatabase.Select(&sendUserInfo,strQuery)
	if conErr != nil {
		return nil,conErr
	}
	if len(sendUserInfo) <= 0 {
		return nil, passwordFailErr
	}
	returnUser := sendUserInfo[0]
	return &returnUser, nil
} 

func TokenGenerator(user *types.SendUserInfo) (*types.TokenData, error){
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