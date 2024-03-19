package usecase

import (
	"enigma-lms/model"
	"enigma-lms/model/dto"
	"enigma-lms/repository"
	"fmt"
	"time"
)

type CourseUseCase interface {
	FindById(id string) (model.Course, error)
	CreateCourse(payload dto.CourseRequestDto) (model.Course, error)
}

type courseUseCase struct {
	repo repository.CourseRepository
}

func (cs *courseUseCase) FindById(id string) (model.Course, error) {
	course, err := cs.repo.Get(id)
	if err != nil {
		return model.Course{}, fmt.Errorf("course with ID %s not found", id)
	}
	return course, nil
}

func (cs *courseUseCase) CreateCourse(payload dto.CourseRequestDto) (model.Course, error) {
	startDate, err := time.Parse("2006-01-02", payload.CourseStartDate)
	if err != nil {
		return model.Course{}, fmt.Errorf("failed to parse CourseStartDate: %s", err.Error())
	}

	endDate, err := time.Parse("2006-01-02", payload.CourseEndDate)
	if err != nil {
		return model.Course{}, fmt.Errorf("failed to parse CourseEndDate: %s", err.Error())
	}

	newCourse := model.Course{
		Id:              payload.Id,
		CourseFullName:  payload.CourseFullName,
		CourseShortName: payload.CourseShortName,
		Description:     payload.Description,
		CourseStartDate: startDate,
		CourseEndDate:   endDate,
		CourseImage:     payload.CourseImage,
	}

	course, err := cs.repo.Create(newCourse)
	if err != nil {
		return model.Course{}, fmt.Errorf("failed to create user: %s", err.Error())
	}
	return course, nil
}

func NewCourseUseCase(repo repository.CourseRepository) CourseUseCase {
	return &courseUseCase{
		repo: repo,
	}
}
