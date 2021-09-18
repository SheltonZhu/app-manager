package jwt

import (
	"fmt"
	"testing"
)

func TestGenerateToken(t *testing.T) {
	token, _ := GenerateToken()
	fmt.Println(token)
	fmt.Println(ParseToken(token))
}