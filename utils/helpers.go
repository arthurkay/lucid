package utils

import (
	"log"
	"math/rand"
	"os"
	"runtime"

	"golang.org/x/crypto/bcrypt"
)

// LucidHomeDir is the deafult configuration directory on the syste
func LucidHomeDir() string {
	configDir, err := os.UserConfigDir()
	if err != nil {
		CheckError(err)
	}
	return configDir + "/lucid"
}

// AppPath returns a tring of the current application path
var AppPath string = getAppPath()

func getAppPath() string {
	path, err := os.Getwd()
	if err != nil {
		CheckError(err)
	}
	return path
}

// CheckError takes in an error value and handles it by logging its output
func CheckError(err error) {
	if err != nil {
		_, fn, line, _ := runtime.Caller(1)
		log.Printf("[error] %s:%d %v", fn, line, err)
	}
}

// CreatePasswordHash creates a bcrypt has with a default cost added
func CreatePasswordHash(password string) ([]byte, bool) {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		CheckError(err)
		return []byte{}, false
	}
	return passwordHash, true

}

// PasswordHashCompare checks whether the password and hash in the database cryptographically match
func PasswordHashCompare(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// RandomString generates a random string of the specififed length
func randomString(n int) string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
