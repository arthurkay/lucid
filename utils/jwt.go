package utils

import (
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
)

var (
	erro error  = godotenv.Load()
	key  string = os.Getenv("AUTH_KEY")
)

func init() {
	if erro != nil {
		panic(erro.Error())
	}
}

type customClaims struct {
	name string
	jwt.StandardClaims
}

// CreateToken generates a jwt string token and an error
func CreateToken() (string, error) {
	claims := customClaims{
		name: "Universe API",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * time.Duration(1)).Unix(),
			Issuer:    "infratel.co.zm",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signature, err := token.SignedString([]byte(key))

	return signature, err
}

// VerifyToken takes in a string token. Then returns a boolean and an error
func VerifyToken(tokenString string) (bool, error) {
	token, err := jwt.ParseWithClaims(tokenString, &customClaims{}, func(token *jwt.Token) (interface{}, error) { return string([]byte(key)), nil })
	if err != nil {
		return false, err
	}
	_, ok := token.Claims.(*customClaims)
	if ok {
		return true, nil
	}
	return true, fmt.Errorf("couldn't verify token")
}
