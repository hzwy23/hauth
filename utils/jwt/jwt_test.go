package jwt

import (
	"fmt"
	"testing"

	jwt "github.com/dgrijalva/jwt-go"
)

var token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE0OTk5MzcwNjcsImlzcyI6Imh6d3kyMyIsIlVzZXJJZCI6ImNhYWRtaW4iLCJEb21haW5JZCI6Im1hcyIsIk9yZ1VuaXRJZCI6Im1hc19qb2luXzM0MTI0IiwiYXV0aG9yaXRpZXMiOiJST0xFX0FETUlOLEFVVEhfV1JJVEUsQUNUVUFUT1IifQ.-xtxhlSyhQjPlCJV1rGFhRm1Ac4_PjpxFNnB8kp7Xjg"

func TestToken(t *testing.T) {
	var jclaim = &JwtClaims{}
	_, err := jwt.ParseWithClaims(token, jclaim, func(*jwt.Token) (interface{}, error) {
		return key, nil
	})
	fmt.Println(err)
}
