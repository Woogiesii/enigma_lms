package middleware

import (
	"enigma-lms/config"
	"enigma-lms/utils/common"
	"net/http"

	"github.com/gin-gonic/gin"
)

func BasicAuth(apiConfig config.ApiConfig) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user, password, ok := ctx.Request.BasicAuth()
		if !ok {
			common.SendErrorResponse(ctx, http.StatusUnauthorized, "Invalid Token")
			return
		}

		if user != apiConfig.ClientId || password != apiConfig.ClientSecret {
			common.SendErrorResponse(ctx, http.StatusUnauthorized, "Invalid Credential")
			return
		}
		ctx.Next()
	}
}
