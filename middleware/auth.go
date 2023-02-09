package middleware

import (
	"golang.org/x/crypto/bcrypt"
	"log"
)

func HashPwds(pwd []byte) string {

	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
}
func ComparePwds(og, obtained string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(og), []byte(obtained))
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
