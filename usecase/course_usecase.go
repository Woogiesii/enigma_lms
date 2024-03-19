package usecase

import (
	"enigma-lms/model"
	"enigma-lms/model/dto"
	"enigma-lms/repository"
	"fmt"
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
	newCourse := model.Course{
		Id:              payload.Id,
		CourseFullName:  payload.CourseFullName,
		CourseShortName: payload.CourseShortName,
		Description:     payload.Description,
		CourseStartDate: payload.CourseStartDate,
		CourseEndDate:   payload.CourseEndDate,
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
