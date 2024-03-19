package controller

import (
	"enigma-lms/model/dto"
	"enigma-lms/usecase"
	"enigma-lms/utils/common"
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
		common.SendErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	common.SendCreateResponse(ctx, "OK", response)
}

func (e *UserController) createHandler(ctx *gin.Context) {
	var payload dto.UserRequestDto
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		common.SendErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	payloadResponse, err := e.uc.CreateUser(payload)

	if err != nil {
		common.SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	common.SendCreateResponse(ctx, "OK", payloadResponse)
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
