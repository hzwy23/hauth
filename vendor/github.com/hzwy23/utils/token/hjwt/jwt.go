package hjwt

import (
	"errors"
	"fmt"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/hzwy23/utils/logs"
)

type JwtClaims struct {
	*jwt.StandardClaims
	User_id   string
	Domain_id string
	Org_id    string
}

var (
	key []byte = []byte("hzwy23@163.com-jwt")
)

func GenToken(user_id, domain_id, org_id string, dt int64) string {

	claims := JwtClaims{
		&jwt.StandardClaims{
			NotBefore: time.Now().Unix(),
			ExpiresAt: time.Now().Unix() + dt,
			Issuer:    "hzwy23",
		},
		user_id,
		domain_id,
		org_id,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(key)
	if err != nil {
		logs.Error(err)
		return ""
	}
	return ss
}

func DestoryToken() string {

	claims := JwtClaims{
		&jwt.StandardClaims{
			NotBefore: int64(time.Now().Unix() - 99998),
			ExpiresAt: int64(time.Now().Unix() - 99999),
			Issuer:    "hzwy23",
		},
		"exit",
		"exit",
		"exit",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(key)
	if err != nil {
		logs.Error(err)
		return ""
	}
	return ss
}

func CheckToken(token string) bool {
	_, err := jwt.Parse(token, func(*jwt.Token) (interface{}, error) {
		return key, nil
	})
	if err != nil {
		fmt.Println("parase with claims failed.", err)
		return false
	}
	return true
}

func ParseJwt(token string) (*JwtClaims, error) {
	var jclaim = &JwtClaims{}
	_, err := jwt.ParseWithClaims(token, jclaim, func(*jwt.Token) (interface{}, error) {
		return key, nil
	})
	if err != nil {
		fmt.Println("parase with claims failed.", err)
		return nil, errors.New("parase with claims failed.")
	}
	return jclaim, nil
}
