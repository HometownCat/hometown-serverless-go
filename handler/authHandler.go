package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"sync"
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
		strQuery += "AND `password` = \"" + *password + "\""
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
	accessTokenClaims := types.AuthTokenData{
		TokenUUID: uuid.NewString(),
		UserUUID: uuid.NewString(),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: jwt.At(time.Now().Add(time.Minute * 30)),
		},
	}
	revokeTokenClamis := types.AuthTokenData{
		TokenUUID: uuid.NewString(),
		UserUUID: uuid.NewString(),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: jwt.At(time.Now().Add(time.Hour * 24 * 30)),
		},
	}

	userBin,_ := json.Marshal(*user)

	var wg sync.WaitGroup

	wg.Add(2)

	go func(){
		json.Unmarshal(userBin,&accessTokenClaims)
		defer wg.Done()
	}()
	go func(){
		json.Unmarshal(userBin,&revokeTokenClamis)
		defer wg.Done()
	}()

	wg.Wait()

	accessToken  := jwt.NewWithClaims(jwt.SigningMethodHS256, &accessTokenClaims)
	revokeToken  := jwt.NewWithClaims(jwt.SigningMethodHS256, &accessTokenClaims)

	accessSignedToken,accessErr := accessToken.SignedString([]byte(os.Getenv("JWT_ACCESS_SECRET_KEY")))	
	revokeSignedToken,revokeErr := revokeToken.SignedString([]byte(os.Getenv("JWT_REVOKE_SECRET_KEY")))	
	if accessErr != nil || revokeErr != nil {
		fmt.Println("access :" +accessErr.Error())
		fmt.Println("reovke :" +revokeErr.Error())
		var newErr error
		if accessErr != nil {
			newErr = accessErr
		}else if revokeErr != nil {
			newErr = revokeErr
		}
		return nil, newErr
	}

	return &types.TokenData{
		AccessToken: accessSignedToken,
		RevokeToken: revokeSignedToken,
	}, nil
}

func TokenParser(token *string, secretKey *string) (*types.AuthTokenData,error){
	returnData := types.AuthTokenData{}

	key := func (token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(*secretKey), nil
	}
	
	tok,err := jwt.ParseWithClaims(*token,&returnData,key)

	if err != nil {
		return nil,err
	}

	if !tok.Valid {
		return nil,errors.New("invalid token")
	}
	return &returnData,nil
}