package controller

import (
	"enigma-lms/model/dto"
	"enigma-lms/usecase"
	"enigma-lms/utils/common"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CourseController struct {
	uc usecase.CourseUseCase
	rg *gin.RouterGroup
}

func (cse *CourseController) getHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	response, err := cse.uc.FindById(id)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	common.SendCreateResponse(ctx, "OK", response)
}

func (cse *CourseController) createHandler(ctx *gin.Context) {
	var payload dto.CourseRequestDto
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		common.SendErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	payloadResponse, err := cse.uc.CreateCourse(payload)

	if err != nil {
		common.SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	common.SendCreateResponse(ctx, "OK", payloadResponse)
}

func (cse *CourseController) Route() {
	cse.rg.GET("/courses/:id", cse.getHandler)
	cse.rg.POST("/courses", cse.createHandler)
}

func NewCourseController(uc usecase.CourseUseCase, rg *gin.RouterGroup) *CourseController {
	return &CourseController{
		uc: uc,
		rg: rg,
	}
}
