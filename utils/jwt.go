package utils

import (
	"fmt"
	"time"

	"github.com/arthurkay/lucid/db"
	"github.com/dgrijalva/jwt-go"
)

var (
	DB         = db.DB{DBPath: LucidHomeDir() + "/lucid.db"}
	key string = DB.Get([]byte("jwt_key"))
)

type customClaims struct {
	name string
	jwt.StandardClaims
}

// CreateToken generates a jwt string token and an error
func CreateToken() (string, error) {
	// Create DB and initialise it with a value
	e := DB.Put([]byte("jwt_key"), []byte(randomString(21)))
	if e != nil {
		return "", e
	}
	claims := customClaims{
		name: "Lucid API",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * time.Duration(1)).Unix(),
			Issuer:    "Arthur Kay",
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
