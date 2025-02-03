package utils

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
)

const (
	key = "VerySecretKey"
	exp = time.Hour
)

func LunaCheck(order int) bool {
	return false
}

func CreateJWT(uId int) (string, error) {
	mySigningKey := []byte(key)

	claims := &jwt.StandardClaims{
		ExpiresAt: time.Now().Unix() + int64(exp),
		Issuer:    fmt.Sprint(uId),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySigningKey)

	if err != nil {
		return "", err
	}
	return ss, nil
}

type MyClaims struct {
	Iss string `json:"iss"`
	jwt.StandardClaims
}

func ValidateJWT(tokenStr string) (int, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(key), nil
	})

	if err != nil {
		return 0, err
	}

	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
		return strconv.Atoi(claims.Iss)
	}

	return 0, errors.New("invalid token")
}
