package jwt_services

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type ApexClaims struct {
	Role string `json:"role"`
	jwt.RegisteredClaims
}

func GenerateJWT(id uint, role string) (string, error) {
	secret := []byte(os.Getenv("TOKEN_SECRET"))

	claims := ApexClaims{
		role,
		jwt.RegisteredClaims{
			Issuer:    "",
			Subject:   strconv.Itoa(int(id)),
			Audience:  []string{},
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			NotBefore: jwt.NewNumericDate(time.Now()),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ID:        "",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, err := token.SignedString(secret)

	return signed, err
}

func ParseJWT(tokenString string, claims ApexClaims, keyFunc jwt.Keyfunc) (ApexClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &ApexClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("AllYourBase"), nil
	})

	if claims, ok := token.Claims.(*ApexClaims); ok && token.Valid {
		fmt.Printf("%v %v", claims.Role, claims.RegisteredClaims.Issuer)
	} else {
		fmt.Println(err)
	}

	return claims, err
}
