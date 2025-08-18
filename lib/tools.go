package lib

import "golang.org/x/crypto/bcrypt"

func HashedPassword(password string)(string) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password),bcrypt.DefaultCost)
	if err != nil{
		return err.Error()
	}

	return string(bytes)
}

func CheckPasswordHash(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash),[]byte(password))
	return err == nil
}