package system

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const OAuthBearer = "Bearer"

type Claims struct {
	jwt.RegisteredClaims
	UserId, RoleId int
}

func GetClaims(tokenStr string) (*Claims, error) {
	tokenStr = tokenStr[strings.Index(tokenStr, " "):]
	var claims Claims
	token, err := jwt.ParseWithClaims(tokenStr, &claims, func(t *jwt.Token) (interface{}, error) {
		fmt.Println("Header::", t.Header)
		return nil, nil
	})
	if err != nil {
		return nil, errors.New("")
	}
	if !token.Valid {
		return nil, errors.New("")
	}
	if claims.ExpiresAt.Before(time.Now().Local()) {
		return nil, errors.New("")
	}
	return &claims, nil
}
