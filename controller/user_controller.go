package controller

import (
	"enigma-lms/model/dto"
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

func (e *UserController) createHandler(ctx *gin.Context) {
	var payload dto.UserRequestDto
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":        http.StatusBadRequest,
			"description": err.Error(),
		})
		return
	}
	payloadResponse, err := e.uc.CreateUser(payload)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":        http.StatusInternalServerError,
			"description": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"code":        http.StatusCreated,
		"description": "OK",
		"data":        payloadResponse,
	})
}

func (e *UserController) getAllUsersHandler(ctx *gin.Context) {
	users, err := e.uc.GetAllUsers()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":        http.StatusInternalServerError,
			"description": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code":        http.StatusOK,
		"description": "OK",
		"data":        users,
	})
}

func (e *UserController) Route() {
	e.rg.GET("/users/:id", e.getHandler)
	e.rg.POST("/users", e.createHandler)
	e.rg.GET("/users", e.getAllUsersHandler)
}

func NewUserController(uc usecase.UserUseCase, rg *gin.RouterGroup) *UserController {
	return &UserController{
		uc: uc,
		rg: rg,
	}
}
