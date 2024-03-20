package common

import (
	"enigma-lms/model"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
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

func JWTAuth(roles ...string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if !strings.Contains(authHeader, "Bearer") {
			SendErrorResponse(ctx, http.StatusForbidden, "Invalid Token")
			return
		}

		jwtSignatureKey := []byte(os.Getenv("SIGNATURE_KEY"))
		tokenString := strings.Replace(authHeader, "Bearer: ", "", -1)
		claims := &JwtClaim{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtSignatureKey, nil
		})
		if err != nil {
			SendErrorResponse(ctx, http.StatusInternalServerError, "Unauthorized User")
			return
		}

		if !token.Valid {
			SendErrorResponse(ctx, http.StatusUnauthorized, "Unauthorized User")
			return
		}

		expiredAt := claims.ExpiresAt
		if time.Now().Unix() > expiredAt {
			SendErrorResponse(ctx, http.StatusUnauthorized, "Expired Token")
			return
		}

		//validation role
		validRole := false
		if len(roles) > 0 {
			for _, role := range roles {
				if role == claims.UserData.Role {
					validRole = true
					break
				}
			}
		}
		if !validRole {
			SendErrorResponse(ctx, http.StatusForbidden, "You don't have permission")
			return
		}

		ctx.Next()
	}
}
