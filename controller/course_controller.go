package controller

import (
	"enigma-lms/model/dto"
	"enigma-lms/usecase"
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
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":        http.StatusInternalServerError,
			"description": err.Error(),
		})
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code":        http.StatusOK,
		"description": "OK",
		"data":        response,
	})
}

func (cse *CourseController) createHandler(ctx *gin.Context) {
	var payload dto.CourseRequestDto
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":        http.StatusBadRequest,
			"description": err.Error(),
		})
		return
	}
	payloadResponse, err := cse.uc.CreateCourse(payload)

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
