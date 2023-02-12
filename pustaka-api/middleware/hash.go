package middleware

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func GetPwd(pw string) []byte {
	return []byte(pw)
}

func HashAndSalt(pw []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pw, bcrypt.MinCost)
	if err != nil {
		fmt.Println(err)
	}
	return string(hash)
}

func ComparePassword(hastedPw string, plainPw []byte) bool {
	byteHash := []byte(hastedPw)

	err := bcrypt.CompareHashAndPassword(byteHash, plainPw)
	if err != nil {
		fmt.Print(err)
		return false
	}

	return true
}
