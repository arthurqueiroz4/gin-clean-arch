package usecase

import (
	"fmt"
	"gin-clean-arch/domain"
)

type UserUsecase struct {
	userRepository domain.UserRepository
}

func NewUserUsecase(userRepository domain.UserRepository) domain.UserUsecase {
	return &UserUsecase{
		userRepository,
	}
}

func (uc *UserUsecase) Create(user domain.User) (*domain.User, error) {
	if userFound, _ := uc.userRepository.FindByEmail(user.Email); userFound.Name != "" {
		return nil, fmt.Errorf("user %v already exists", user.Name)
	}

	userCreated, err := uc.userRepository.Create(&user)
	if err != nil {
		return nil, err
	}

	return userCreated, err
}

func (uc *UserUsecase) FindByEmail(email string) (*domain.User, error) {
	userFound, err := uc.userRepository.FindByEmail(email)
	if err != nil {
		return nil, err
	}
	if userFound.Email == "" {
		return nil, fmt.Errorf("user not found")
	}
	return userFound, nil
}
