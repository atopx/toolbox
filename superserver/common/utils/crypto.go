package utils

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/golang-jwt/jwt/v5"
	"strings"
	"time"
)

func Hash(value string) string {
	if value == "" {
		return value
	}
	buf := []byte(value)
	m := md5.New()
	m.Write(buf)
	return hex.EncodeToString(m.Sum(buf[0:1]))
}

func UnSignTokenExpire(tokenStr string, key any) (expire *jwt.NumericDate, err error) {
	tokenStr = tokenStr[strings.Index(tokenStr, " ")+1:]
	var claims jwt.RegisteredClaims
	_, err = jwt.ParseWithClaims(tokenStr, &claims, func(t *jwt.Token) (interface{}, error) {
		return key, nil
	})
	if err != nil {
		return nil, err
	}
	return claims.ExpiresAt, nil
}

func SignToken(issued, expires time.Time, key any) string {
	claims := jwt.RegisteredClaims{
		IssuedAt:  &jwt.NumericDate{Time: issued},
		ExpiresAt: &jwt.NumericDate{Time: expires},
	}
	token, _ := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(key)
	return token
}
