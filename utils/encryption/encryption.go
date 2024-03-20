package encryption

import "golang.org/x/crypto/bcrypt"

/*
TODO:
1. Create Function Hash Password
2. Create Function for Compare Password
*/

func HashPassword(pass string) (string, error) {
	passHash, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(passHash), nil
}

func CheckPassword(pass, hashPass string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashPass), []byte(pass))
	return err == nil
}
