package system

import (
	"cloudos/internal/model/user_token"
	"errors"
	"strings"
	"time"

	"github.com/spf13/viper"

	"github.com/golang-jwt/jwt/v5"
)

var secretKey []byte

func init() {
	secretKey = []byte(viper.GetString("admin.secret"))
}

type Claims struct {
	jwt.RegisteredClaims
	UserId    int
	IsRefresh bool
}

func UnSignClaims(tokenStr string) (*Claims, error) {
	tokenStr = tokenStr[strings.Index(tokenStr, " ")+1:]
	var claims Claims
	token, err := jwt.ParseWithClaims(tokenStr, &claims, func(t *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, errors.New("token invalid")
	}
	return &claims, nil
}

func SignClaims(userId int) (token user_token.UserToken) {
	current := time.Now().Local()
	expires := current.Add(24 * time.Hour)
	token.IssuedTime = current.Unix()
	token.ExpireTime = expires.Unix()
	token.AccessToken, _ = sign(&Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  &jwt.NumericDate{Time: current},
			ExpiresAt: &jwt.NumericDate{Time: expires},
		},
		UserId:    userId,
		IsRefresh: false,
	})
	token.RefreshToken, _ = sign(&Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  &jwt.NumericDate{Time: current},
			ExpiresAt: &jwt.NumericDate{Time: current.Add(7 * 24 * time.Hour)},
		},
		UserId:    userId,
		IsRefresh: true,
	})
	return token
}

func sign(claims *Claims) (string, error) {
	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(secretKey)
}
