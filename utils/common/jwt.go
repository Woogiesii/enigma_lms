package common

import (
	"enigma-lms/model"
	"os"

	"github.com/dgrijalva/jwt-go"
)

type JwtClaim struct {
	jwt.StandardClaims
	UserData model.UserData `json:"user"`
}

var (
	appName          = os.Getenv("APP_NAME")
	jwtSigningMethod = jwt.SigningMethodHS256
	jwtSignatureKey  = []byte(os.Getenv("SIGNATURE_KEY"))
)

func GenerateTokenJwt(userData model.User, expiredAt int64) (string, error) {
	claims := JwtClaim{
		StandardClaims: jwt.StandardClaims{
			Issuer:    appName,
			ExpiresAt: expiredAt,
		},
		UserData: model.UserData{
			Id:        userData.Id,
			FirstName: userData.FirstName,
			LastName:  userData.LastName,
			Email:     userData.Email,
			Role:      userData.Role,
			Photo:     userData.Photo,
		},
	}

	token := jwt.NewWithClaims(jwtSigningMethod, claims)
	signedToken, err := token.SignedString(jwtSignatureKey)
	if err != nil {
		return "", err
	}
	return signedToken, nil
}
