package utils

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

var signatureKEY []byte

type DataClaims struct {
	ID    string `json:"id"`
	Email string `json:"email"`
}

type Claims struct {
	DataClaims
	jwt.StandardClaims
}

func InitJWT(key string) {
	signatureKEY = []byte(key)
}

func NewToken(params DataClaims) *Claims {
	return &Claims{
		DataClaims: params,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}
}

func (c *Claims) Create() (string, error) {
	tokens := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	signedStr, err := tokens.SignedString(signatureKEY)
	if err != nil {
		return "", fmt.Errorf("failed to SignedString : %v", err)
	}
	return signedStr, nil
}

func VerifyToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, isValid := token.Method.(*jwt.SigningMethodHMAC); !isValid {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return signatureKEY, nil
	})

	if err != nil {
		return nil, err
	}

	return token, nil
}

func DecodeToken(tokenString string) (jwt.MapClaims, error) {
	token, err := VerifyToken(tokenString)
	if err != nil {
		return nil, err
	}

	claims, isOk := token.Claims.(jwt.MapClaims)
	if isOk && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("invalid token")
}
