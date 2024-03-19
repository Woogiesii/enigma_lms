package usecase

import (
	"enigma-lms/model"
	"enigma-lms/model/dto"
	"enigma-lms/repository"
	"fmt"
)

type UserUseCase interface {
	FindById(id string) (model.User, error)
	CreateUser(payload dto.UserRequestDto) (model.User, error)
	GetAllUsers() ([]model.User, error)
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

func (u *userUseCase) CreateUser(payload dto.UserRequestDto) (model.User, error) {
	newUser := model.User{
		Id:        payload.Id,
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
		Email:     payload.Email,
		Username:  payload.Username,
		Password:  payload.Password,
		Role:      payload.Role,
		Photo:     payload.Photo,
	}

	user, err := u.repo.Create(newUser)
	if err != nil {
		return model.User{}, fmt.Errorf("failed to create user: %s", err.Error())
	}

	return user, nil
}

func (u *userUseCase) GetAllUsers() ([]model.User, error) {
	users, err := u.repo.GetAll()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func NewUserUseCase(repo repository.UserRepository) UserUseCase {
	return &userUseCase{
		repo: repo,
	}
}
