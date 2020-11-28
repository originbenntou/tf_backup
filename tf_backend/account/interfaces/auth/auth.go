package auth

import (
	"fmt"
	"os"
	"time"

	"github.com/TrendFindProject/tf_backend/account/constant"
	"github.com/dgrijalva/jwt-go"
	"github.com/sethvargo/go-password/password"
)

func GenerateNewTokenByUuid(uid string) string {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["admin"] = true
	claims["user"] = uid
	claims["iat"] = time.Now()
	claims["exp"] = time.Now().Add(time.Hour * constant.JwtExpire).Unix()

	// signature
	signed, _ := token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))

	return signed
}

func ValidateTokenString(ts string) (*jwt.Token, error) {
	token, err := jwt.Parse(ts, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("JWT_SECRET_KEY")), nil
	})
	if err != nil {
		return nil, err
	}

	return token, nil
}

func GetUserUuidFromClaim(t *jwt.Token) string {
	claims, ok := t.Claims.(jwt.MapClaims)
	if !ok {
		return ""
	}

	uc, ok := claims["user"].(string)
	if !ok {
		return ""
	}

	return uc
}

func GeneratePassword() string {
	p, err := password.Generate(16, 0, 0, false, false)
	if err != nil {
		return ""
	}

	return p
}
