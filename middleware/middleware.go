package middleware

import (
	"enigma-lms/utils/common"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func BasicAuth(ctx *gin.Context) {
	user, password, ok := ctx.Request.BasicAuth()
	if !ok {
		common.SendErrorResponse(ctx, http.StatusUnauthorized, "Invalid Token")
		return
	}

	if user != os.Getenv("CLIENT_ID") || password != os.Getenv("CLIENT_SECRET") {
		common.SendErrorResponse(ctx, http.StatusUnauthorized, "Invalid Credential")
	}
	ctx.Next()
}
