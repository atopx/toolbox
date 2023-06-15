package utils

import (
	"crypto/md5"
	crand "crypto/rand"
	"encoding/binary"
	"encoding/hex"
	"math/big"
	"math/rand"
	"strings"
	"sync"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var seed int64

func Hash(value string) string {
	if value == "" {
		return value
	}
	buf := []byte(value)
	m := md5.New()
	m.Write(buf)
	return hex.EncodeToString(m.Sum(buf[0:1]))
}

func UnSignTokenExpire(tokenStr string, key int64) (expire *jwt.NumericDate, err error) {
	tokenStr = tokenStr[strings.Index(tokenStr, " ")+1:]
	var claims jwt.RegisteredClaims
	_, err = jwt.ParseWithClaims(tokenStr, &claims, func(t *jwt.Token) (interface{}, error) {
		return Int64ToBytes(key), nil
	})
	if err != nil {
		return nil, err
	}
	return claims.ExpiresAt, nil
}

func SignToken(issued, expires time.Time, key int64) string {
	claims := jwt.RegisteredClaims{
		IssuedAt:  &jwt.NumericDate{Time: issued},
		ExpiresAt: &jwt.NumericDate{Time: expires},
	}
	token, _ := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(Int64ToBytes(key))
	return token
}

func Int64ToBytes(value int64) []byte {
	return big.NewInt(value).Bytes()
}

func NewTraceId() string {
	var s sync.Mutex
	s.Lock()
	defer s.Unlock()
	_ = binary.Read(crand.Reader, binary.LittleEndian, &seed)
	tid := [16]byte{}
	rand.New(rand.NewSource(seed)).Read(tid[:])
	return hex.EncodeToString(tid[:])
}
