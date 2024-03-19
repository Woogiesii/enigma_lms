package usecase

import (
	"enigma-lms/model"
	"enigma-lms/model/dto"
	"enigma-lms/repository"
	"fmt"
)

type EnrollmentUseCase interface {
	RegisterNewEnrollment(payload dto.EnrollmentRequestDto) (model.Enrollment, error)
}

type enrollmentUseCase struct {
	repo     repository.EnrollmentRepository
	userUC   UserUseCase
	courseUC CourseUseCase
}

func (e *enrollmentUseCase) RegisterNewEnrollment(payload dto.EnrollmentRequestDto) (model.Enrollment, error) {
	course, err := e.courseUC.FindById(payload.CourseId)
	if err != nil {
		return model.Enrollment{}, err
	}

	var newEnrollmentDetail []model.EnrollmentDetail

	for _, v := range payload.Users {
		user, err := e.userUC.FindById(v)
		if err != nil {
			return model.Enrollment{}, err
		}
		newEnrollmentDetail = append(newEnrollmentDetail, model.EnrollmentDetail{User: user})
	}

	newEnrollment := model.Enrollment{
		Course:            course,
		EnrollmentDetails: newEnrollmentDetail,
	}

	enrollment, err := e.repo.Create(newEnrollment)
	if err != nil {
		return model.Enrollment{}, fmt.Errorf("failed to create enrollment: %s", err.Error())
	}
	return enrollment, nil
}

func NewEnrollmentUseCase(repo repository.EnrollmentRepository, userUC UserUseCase, courseUC CourseUseCase) EnrollmentUseCase {
	return &enrollmentUseCase{
		repo:     repo,
		userUC:   userUC,
		courseUC: courseUC,
	}
}
