package manager

import "enigma-lms/usecase"

type UseCaseManager interface {
	UserUseCase() usecase.UserUseCase
	CourseUseCase() usecase.CourseUseCase
	EnrollmentUseCase() usecase.EnrollmentUseCase
}

type useCaseManager struct {
	repo RepoManager
}

func (u *useCaseManager) CourseUseCase() usecase.CourseUseCase {
	return usecase.NewCourseUseCase(u.repo.CourseRepo())
}

func (u *useCaseManager) EnrollmentUseCase() usecase.EnrollmentUseCase {
	return usecase.NewEnrollmentUseCase(u.repo.EnrollRepo(), u.UserUseCase(), u.CourseUseCase())
}

func (u *useCaseManager) UserUseCase() usecase.UserUseCase {
	return usecase.NewUserUseCase(u.repo.UserRepo())
}

func NewUseCaseManager(repo RepoManager) UseCaseManager {
	return &useCaseManager{repo: repo}
}
