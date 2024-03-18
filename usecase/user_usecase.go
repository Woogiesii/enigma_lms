package usecase

import (
	"enigma-lms/model"
	"enigma-lms/repository"
	"fmt"
)

type UserUseCase interface {
	FindById(id string) (model.User, error)
}

type userUseCase struct {
	repo repository.UserRepository
}

func (u *userUseCase) FindById(id string) (model.User, error) {
	user, err := u.repo.Get(id)
	if err != nil {
		return model.User{}, fmt.Errorf("user with ID %s not found", id)
	}
	return user, nil
}

func NewUserUseCase(repo repository.UserRepository) UserUseCase {
	return &userUseCase{
		repo: repo,
	}
}
