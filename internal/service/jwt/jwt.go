package jwt

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	Email  string `json:"email"`
	UserId int64  `json:"userId"`
	jwt.RegisteredClaims
}

func Create(userid int64, email string) (string, error) {
	claims := Claims{
		UserId: userid,
		Email:  email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(36 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	strToken, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))

	return strToken, err
}

func Parse(tokenString string) (*jwt.Token, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET_KEY")), nil
	})
	return token, err
}

func ParseGetClaims(tokenString string) (*Claims, error) {
	token, err := Parse(tokenString)
	return token.Claims.(*Claims), err
}
