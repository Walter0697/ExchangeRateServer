package utility

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func GetEncrpytPassword(pwd string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
}

func ComparePassword(hashedPwd string, plainPwd string) bool {
	byteHash := []byte(hashedPwd)
	comparingPwd := []byte(plainPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, comparingPwd)
	return err == nil
}
