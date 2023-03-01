package system

import (
	"errors"
	"fmt"
	"github.com/spf13/viper"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const OAuthBearer = "Bearer"

type Claims struct {
	jwt.RegisteredClaims
	UserId, RoleId int
}

func UnSignClaims(tokenStr string) (*Claims, error) {
	tokenStr = tokenStr[strings.Index(tokenStr, " ")+1:]
	var claims Claims
	fmt.Println(tokenStr)
	token, err := jwt.ParseWithClaims(tokenStr, &claims, func(t *jwt.Token) (interface{}, error) {
		return loadSecretKey(), nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, errors.New("token invalid")
	}
	return &claims, nil
}

func SignClaims(userId, roleId int) (string, error) {
	currentTime := time.Now().Local()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: &jwt.NumericDate{Time: currentTime.Add(3000 * time.Hour)},
			IssuedAt:  &jwt.NumericDate{Time: currentTime},
		},
		UserId: userId,
		RoleId: roleId,
	})
	tokenStr, err := token.SignedString(loadSecretKey())
	if err != nil {
		return tokenStr, err
	}
	viper.GetString("admin.secret")
	// use the username as secret key
	return fmt.Sprintf("%s %s", OAuthBearer, tokenStr), nil
}

func loadSecretKey() []byte {
	return []byte(viper.GetString("admin.secret"))
}
