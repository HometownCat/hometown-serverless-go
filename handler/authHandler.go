package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go/v4"
	"github.com/google/uuid"
	"hometown.com/hometown-serverless-go/modules/database"
	"hometown.com/hometown-serverless-go/modules/redis"
	"hometown.com/hometown-serverless-go/types"
)

var (
	tables = *database.MysqlTables()
)

func SignUp(user *types.User) error {
	// db 처리 추가
	strQuery := "INSERT INTO `user` SET username = \"" + user.Username + "\", email = \"" + user.Email + "\", password = \"" + user.Password + "\", userIp = \"" + user.UserIp + "\", status = 0"
	if user.Address != nil {
		strQuery += ", address = \"" + *user.Address + "\""
	}
	if user.PhoneNumber != nil {
		strQuery += ", phoneNumber = \"" + *user.PhoneNumber + "\""
	}
	if user.ProfileImage != nil {
		strQuery += ", profileImage = \"" + *user.ProfileImage + "\""
	}

	result, conErr := database.MasterDatabase.Exec(strQuery)

	if conErr == nil {
		id, rowErr := result.LastInsertId()
		conErr = rowErr
		user.Id = &id
	}
	return conErr
}

func GetUser(email *string, password *string, userData *types.SendUserInfo) error {
	// db 처리 추가
	var sendUserInfo []types.SendUserInfo
	var passwordFailErr error
	strQuery := "SELECT id, email, username, address, phoneNumber, profileImage FROM `user` WHERE `email` = \"" + *email + "\""
	if password != nil {
		passwordFailErr = errors.New("password not matched")
		strQuery += "AND `password` = \"" + *password + "\""
	}

	conErr := database.SlaveDatabase.Select(&sendUserInfo, strQuery)
	if conErr != nil {
		return conErr
	}
	if len(sendUserInfo) <= 0 {
		return passwordFailErr
	}

	returnUser := sendUserInfo[0]
	userBin, _ := json.Marshal(returnUser)
	json.Unmarshal(userBin, &userData)
	return nil
}

func SetUserToken(accessToken *string, revokeToken *string, id *int64) error {
	strQuery := "INSERT INTO `" +
		tables["userToken"] +
		"` SET userId = ?, accessToken = ?, revokeToken = ?" +
		" ON DUPLICATE KEY UPDATE accessToken = ?, revokeToken = ?"

	_, conErr := database.MasterDatabase.Exec(
		strQuery,
		[]interface{}{*id, *accessToken, *revokeToken, *accessToken, *revokeToken}...,
	)

	return conErr
}

func RedisTokenGenerator(user *types.SendUserInfo, ) error {
	userBin, _ := json.Marshal(*user)

	expired := time.Hour * 24 * 30
	token := uuid.NewString()
	valueStr := string(userBin)
	setErr := redis.SetData(&token, &valueStr, &expired)
	if setErr != nil {
		return setErr
	}
	user.AccessToken = &token
	return nil
}

func TokenGenerator(user *types.SendUserInfo, secretKey *string, avaliableTime *time.Duration) (*string, error) {
	tokenClaims := types.AuthTokenData{
		TokenUUID: uuid.NewString(),
		UserUUID:  uuid.NewString(),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: jwt.At(time.Now().Add(*avaliableTime)),
		},
	}

	userBin, _ := json.Marshal(*user)

	json.Unmarshal(userBin, &tokenClaims)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims)
	signedToken, err := token.SignedString([]byte(*secretKey))

	if err != nil {
		return nil, err
	}
	return &signedToken, nil
}

func TokenParser(token *string, userInfo *types.SendUserInfo) error {
	// tokenData := types.AuthTokenData{}
	// key := func(token *jwt.Token) (interface{}, error) {
	// 	if _, isValid := token.Method.(*jwt.SigningMethodHMAC); !isValid {
	// 		return nil, errors.New("unexpected signing method")
	// 	}
	// 	return []byte(*secretKey), nil
	// }

	// tok, err := jwt.ParseWithClaims(*token, &tokenData, key)

	// if err != nil {
	// 	return nil, err
	// }
	// if !tok.Valid {
	// 	return nil, errors.New("invalid token")
	// }

	// dataBin, _ := json.Marshal(tokenData)

	redisType := "string"

	userData, getErr := redis.GetData(token, &redisType)

	userStr := fmt.Sprintf("%v", *userData)

	if getErr != nil {
		return getErr
	}
	json.Unmarshal([]byte(userStr), userInfo)
	return nil
}
