package usecase

import (
	"fmt"
	"gin-clean-arch/domain"
	"strconv"
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

func (uc *UserUsecase) FindById(id uint) (*domain.User, error) {
	userFound, err := uc.userRepository.FindById(id)
	if err != nil {
		return nil, err
	}
	if userFound.Name == "" {
		return nil, fmt.Errorf("user with id equal %v not found", err.Error())
	}
	return userFound, nil
}

func (uc *UserUsecase) FindAll(page, pageSize string) (*[]domain.User, int64, error) {
	pageInt, err := strconv.Atoi(page)
	if err != nil {
		return nil, 0, err
	}
	sizeInt, err := strconv.Atoi(pageSize)
	if err != nil {
		return nil, 0, err
	}
	allUsers, total, err := uc.userRepository.FindAll(pageInt, sizeInt)
	if err != nil {
		return nil, 0, err
	}
	return allUsers, total, nil
}

func (uc *UserUsecase) Delete(id int) error {
	return uc.userRepository.Delete(uint(id))
}
