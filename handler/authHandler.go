package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go/v4"
	"github.com/google/uuid"
	"hometown.com/hometown-serverless-go/modules/database"
	"hometown.com/hometown-serverless-go/types"
)


func SignUp(user *types.User) error {
	// db 처리 추가
	strQuery := "INSERT INTO `user` SET username = \"" + user.Username +  "\", email = \"" + user.Email + "\", password = \"" + user.Password + "\", userIp = \"" + user.UserIp +"\", status = 0"
	if user.Address != nil {
		strQuery += ", address = \"" + *user.Address + "\""
	}
	if user.PhoneNumber != nil {
		strQuery += ", phoneNumber = \"" + *user.PhoneNumber + "\""
	}
	if user.ProfileImage != nil {
		strQuery += ", profileImage = \"" + *user.ProfileImage + "\""
	}

	result,conErr := database.MasterDatabase.Exec(strQuery)

	if conErr == nil {
		id,rowErr := result.LastInsertId()
		conErr = rowErr
		user.Id = &id
	}
	return conErr
}

func GetUser(email *string , password *string) (*types.SendUserInfo, error){
	// db 처리 추가
	var sendUserInfo []types.SendUserInfo
	var passwordFailErr error
	strQuery := "SELECT id, email, username, address, phoneNumber, profileImage FROM `user` WHERE `email` = \"" + *email + "\""
	if password != nil {
		passwordFailErr = errors.New("password not matched");
		strQuery += "AND `password` = \"" + *password + "\""
	}
	fmt.Println(strQuery)
	conErr := database.SlaveDatabase.Select(&sendUserInfo,strQuery)
	if conErr != nil {
		return nil,conErr
	}
	if len(sendUserInfo) <= 0 {
		return nil, passwordFailErr
	}
	returnUser := sendUserInfo[0]
	return &returnUser, nil
} 

func TokenGenerator(user *types.SendUserInfo, secretKey *string, avaliableTime *time.Duration) (*string, error){
	tokenClaims := types.AuthTokenData{
		TokenUUID: uuid.NewString(),
		UserUUID: uuid.NewString(),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: jwt.At(time.Now().Add(*avaliableTime)),
		},
	}

	userBin,_ := json.Marshal(*user)

	json.Unmarshal(userBin,&tokenClaims)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims)
	signedToken,err := token.SignedString([]byte(*secretKey))	

	if err != nil {
		return nil, err
	}
	return &signedToken,nil
}

func TokenParser(token *string, secretKey *string) (*types.SendUserInfo,error){
	tokenData := types.AuthTokenData{}
	returnData := types.SendUserInfo{}
	key := func (token *jwt.Token) (interface{}, error) {
		if _, isValid := token.Method.(*jwt.SigningMethodHMAC); !isValid {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(*secretKey), nil
	}
	
	tok,err := jwt.ParseWithClaims(*token,&tokenData,key)

	if err != nil {
		return nil,err
	}
	if !tok.Valid {
		return nil,errors.New("invalid token")
	}

	dataBin,_ := json.Marshal(tokenData)

	json.Unmarshal(dataBin,&returnData);

	return &returnData,nil
}