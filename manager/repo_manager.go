package manager

import "enigma-lms/repository"

type RepoManager interface {
	UserRepo() repository.UserRepository
	CourseRepo() repository.CourseRepository
	EnrollRepo() repository.EnrollmentRepository
}

type repoManager struct {
	infra InfraManager
}

func (r *repoManager) CourseRepo() repository.CourseRepository {
	return repository.NewCourseRepository(r.infra.Conn())
}

func (r *repoManager) EnrollRepo() repository.EnrollmentRepository {
	return repository.NewEnrollmentRepository(r.infra.Conn())
}

func (r *repoManager) UserRepo() repository.UserRepository {
	return repository.NewUserRepository(r.infra.Conn())
}

func NewRepoManager(infra InfraManager) RepoManager {
	return &repoManager{infra: infra}
}
