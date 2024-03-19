package controller

import (
	"enigma-lms/model/dto"
	"enigma-lms/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type EnrollmentController struct {
	uc usecase.EnrollmentUseCase
	rg *gin.RouterGroup
}

func (e *EnrollmentController) createHandler(ctx *gin.Context) {
	var payload dto.EnrollmentRequestDto
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":        http.StatusBadRequest,
			"description": err.Error(),
		})
		return
	}
	payloadResponse, err := e.uc.RegisterNewEnrollment(payload)

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

func (e *EnrollmentController) Route() {
	e.rg.POST("/enrollments", e.createHandler)
}

func NewEnrollmentController(uc usecase.EnrollmentUseCase, rg *gin.RouterGroup) *EnrollmentController {
	return &EnrollmentController{
		uc: uc,
		rg: rg,
	}
}
