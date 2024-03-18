package controller

import (
	"enigma-lms/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	uc usecase.UserUseCase
	rg *gin.RouterGroup
}

func (e *UserController) getHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	response, err := e.uc.FindById(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Code":        http.StatusInternalServerError,
			"description": err.Error(),
		})
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code":        http.StatusOK,
		"description": "OK",
		"data":        response,
	})
}

func (e *UserController) Route() {
	e.rg.GET("/users/:id", e.getHandler)
}

func NewUserController(uc usecase.UserUseCase, rg *gin.RouterGroup) *UserController {
	return &UserController{
		uc: uc,
		rg: rg,
	}
}
