package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

type Claims struct {
	ID       int64
	Username string
	jwt.StandardClaims
}

const secret = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9"

func GenerateToken() (string, error) {
	duration := 300 * time.Second
	nowTime := time.Now()
	expireTime := nowTime.Add(duration)
	issuer := "Shelton"
	claims := Claims{
		ID:       10086,
		Username: "Shelton",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    issuer,
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(secret))
	return token, err
}

func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
