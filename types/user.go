package types

import "time"

type User struct {
	Id *int64 `json:"id" db:"id"`
	Username string `json:"username" db:"username"`
	Email string `json:"email" db:"email"`
	Password string `json:"password" db:"password"`
	Address *string `json:"address" db:"address"`
	PhoneNumber *string `json:"phoneNumber" db:"phoneNumber"`
	ProfileImage *string `json:"profileImage" db:"profileImage"`
	UserIp string `json:"userIp" db:"userIp"`
	Status *int8 `json:"status" db:"status"`
	CreatedAt time.Time `json:"createdAt" db:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt" db:"updatedAt"`
}

type SendUserInfo struct {
	Id int64 `json:"id" db:"id"`
	Email string `json:"email" db:"email"`
	Username string `json:"name" db:"name"`
	Password string `json:"password" db:"password"`
	Address *string `json:"address" db:"address"`
	PhoneNumber *string `json:"phoneNumber" db:"phoneNumber"`
	ProfileImage *string `json:"profileImage" db:"profileImage"`
}

type TokenData struct{
	AccessToken string `json:"accesstoken"`
	RevokeToken string `json:"revokeToken"`
}

type LoginUser struct {
	Email string `json:"email"`
	Password string `json:"password"`
}