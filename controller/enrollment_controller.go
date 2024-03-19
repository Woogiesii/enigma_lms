package controller

import (
	"enigma-lms/model/dto"
	"enigma-lms/usecase"
	"enigma-lms/utils/common"
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
		common.SendErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	payloadResponse, err := e.uc.RegisterNewEnrollment(payload)

	if err != nil {
		common.SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	common.SendCreateResponse(ctx, "OK", payloadResponse)
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
