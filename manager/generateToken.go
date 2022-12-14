package manager

import (
	"crypto/sha256"
	"encoding/hex"

	"hometown.com/hometown-serverless-go/handler"
	"hometown.com/hometown-serverless-go/types"
)

// var JWT_ACCESS_SECRET_KEY = os.Getenv("JWT_ACCESS_SECRET_KEY")
// var JWT_REVOKE_SECRET_KEY = os.Getenv("JWT_REVOKE_SECRET_KEY")

// var JWT_ACCESS_AVALIABLE_TIME = time.Minute * 30

// var JWT_REVOKE_AVALIABLE_TIME = time.Hour * 24 * 30

func TokenGenerator(email *string, password *string, userInfo *types.SendUserInfo) error {
	hashPassword := sha256.Sum256([]byte(*password))
	*password = hex.EncodeToString(hashPassword[:])

	getUserErr := handler.GetUser(email, password, userInfo)
	if getUserErr != nil {
		return getUserErr
	}

	// tokenData := types.TokenData{}
	// var tokenErr error
	// var wg sync.WaitGroup

	// wg.Add(2)

	// go func() {
	// 	accessToken, accessErr := handler.TokenGenerator(&userData, &JWT_ACCESS_SECRET_KEY, &JWT_ACCESS_AVALIABLE_TIME)
	// 	tokenData.AccessToken = *accessToken
	// 	if accessErr != nil {
	// 		tokenErr = accessErr
	// 	}
	// 	defer wg.Done()
	// }()

	// go func() {
	// 	revokeToken, revokeErr := handler.TokenGenerator(&userData, &JWT_ACCESS_SECRET_KEY, &JWT_ACCESS_AVALIABLE_TIME)
	// 	tokenData.RevokeToken = *revokeToken
	// 	if revokeErr != nil {
	// 		tokenErr = revokeErr
	// 	}
	// 	defer wg.Done()
	// }()

	// wg.Wait()

	// if tokenErr != nil {
	// 	return nil, tokenErr
	// }

	// tokenBin, _ := json.Marshal(tokenData)

	// json.Unmarshal(tokenBin, &userData)
	tokenErr := handler.RedisTokenGenerator(userInfo)

	if tokenErr != nil {
		return tokenErr
	}

	return nil
}
