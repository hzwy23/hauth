package hjwt

import (
	"fmt"
	"testing"
)

func TestToken(t *testing.T) {
	token := GenToken()
	fmt.Println(token)
	flag := CheckToken(token)
	fmt.Println(flag)
}
