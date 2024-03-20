package controller

import (
	"enigma-lms/config"
	"enigma-lms/middleware"
	"enigma-lms/model/dto"
	"enigma-lms/usecase"
	"enigma-lms/utils/common"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	uc     usecase.UserUseCase
	rg     *gin.RouterGroup
	apiCfg config.ApiConfig
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

func (e *UserController) loginHandler(ctx *gin.Context) {
	var payload dto.LoginRequestDto
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		common.SendErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	loginData, err := e.uc.LoginUser(payload)
	if err != nil {
		if err.Error() == "1" {
			common.SendErrorResponse(ctx, http.StatusForbidden, "Invalid Password")
			return
		}
		common.SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	fmt.Println("Login Data:", loginData)
	common.SendSingleResponse(ctx, "OK", loginData)
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
		common.SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	common.SendCreateResponse(ctx, "OK", users)
}

func (e *UserController) Route() {
	userGroup := e.rg.Group("/users")
	{
		userGroup.GET("/:id", e.getHandler) // GET BY ID
		userGroup.POST("", e.createHandler) //CREATE USERS
		userGroup.GET("", e.getAllUsersHandler)
		userGroup.POST("/login", middleware.BasicAuth(e.apiCfg), e.loginHandler) //GET ALL
	}
}

func NewUserController(uc usecase.UserUseCase, rg *gin.RouterGroup, apiCfg config.ApiConfig) *UserController {
	return &UserController{
		uc:     uc,
		rg:     rg,
		apiCfg: apiCfg,
	}
}
