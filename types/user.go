package types

type User struct {
	Email string `json:"email" db:"email"`
	Password string `json:"password" db:"password"`
	Name string `json:"name" db:"name"`
}

type TokenData struct{
	AccessToken string `json:"accesstoken"`
	RevokeToken string `json:"revokeToken"`
}