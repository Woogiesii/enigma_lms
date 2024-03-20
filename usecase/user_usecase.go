package usecase

import (
	"enigma-lms/model"
	"enigma-lms/model/dto"
	"enigma-lms/repository"
	"enigma-lms/utils/common"
	"enigma-lms/utils/encryption"
	"errors"
	"fmt"
	"time"
)

type UserUseCase interface {
	FindById(id string) (model.User, error)
	CreateUser(payload dto.UserRequestDto) (model.User, error)
	GetAllUsers() ([]model.User, error)
	LoginUser(in dto.LoginRequestDto) (dto.LoginResponseDto, error)
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

func (u *userUseCase) LoginUser(in dto.LoginRequestDto) (dto.LoginResponseDto, error) {
	// get data from user by username
	userData, err := u.repo.GetByUsername(in.Username)
	if err != nil {
		return dto.LoginResponseDto{}, err
	}
	// compare password from user and from request
	isValid := encryption.CheckPassword(in.Pass, userData.Password)
	if !isValid {
		return dto.LoginResponseDto{}, errors.New("1")
	}
	// Generate token JWT
	loginExpDuration := time.Duration(10) * time.Minute
	expiresAt := time.Now().Add(loginExpDuration).Unix()
	accesToken, err := common.GenerateTokenJwt(userData, expiresAt)
	if err != nil {
		return dto.LoginResponseDto{}, err
	}

	return dto.LoginResponseDto{
		AccesToken: accesToken,
		UserId:     userData.Id,
	}, nil
}

func (u *userUseCase) CreateUser(payload dto.UserRequestDto) (model.User, error) {
	hashPassword, err := encryption.HashPassword(payload.Password)
	if err != nil {
		return model.User{}, nil
	}
	newUser := model.User{
		Id:        payload.Id,
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
		Email:     payload.Email,
		Username:  payload.Username,
		Password:  hashPassword,
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
