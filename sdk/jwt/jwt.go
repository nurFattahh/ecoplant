package crypto

import (
	"ecoplant/entity"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func GenerateToken(payload entity.User) (string, error) {
	expStr := os.Getenv("JWT_EXP")
	var exp time.Duration
	exp, err := time.ParseDuration(expStr)
	if expStr == "" || err != nil {
		exp = time.Hour * 1
	}

	temporaryJwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       payload.ID,
		"username": payload.Username,
		"exp":      time.Now().Add(exp).Unix(),
	})

	tokenJwt, err := temporaryJwtToken.SignedString([]byte(os.Getenv("secret_key")))
	if err != nil {
		return "", err
	}
	return tokenJwt, nil
}
