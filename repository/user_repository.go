package repository

import (
	"fmt"
	"gin-clean-arch/domain"
	"gorm.io/gorm"
)

type userRepository struct {
	database *gorm.DB
}

func NewUserRepository(db *gorm.DB) domain.UserRepository {
	return &userRepository{
		database: db,
	}
}

func (ur *userRepository) Create(user *domain.User) (*domain.User, error) {
	userFound, _ := ur.FindByEmail(user.Email)

	if userFound.Email != "" {
		return nil, fmt.Errorf("email already registered")
	}

	err := ur.database.Create(&user).Error

	if err != nil {
		return nil, fmt.Errorf("failed create user %v", err.Error())
	}

	return user, nil
}

func (ur *userRepository) FindById(id uint) (*domain.User, error) {
	var person = domain.User{}

	err := ur.database.Where("id = ?", id).First(&person).Error

	if err != nil {
		return nil, fmt.Errorf("id not exists")
	}

	return &person, nil
}

func (ur *userRepository) FindByEmail(email string) (*domain.User, error) {
	var person domain.User

	err := ur.database.Where("email = ?", email).Find(&person).Error

	if err != nil {
		return nil, fmt.Errorf("email not exists")
	}

	return &person, nil
}

func (ur *userRepository) FindAll(page, pageSize int) (*[]domain.User, int64, error) {
	var persons []domain.User

	offset := (page - 1) * pageSize

	db := ur.database.Select("id", "name", "email", "pass", "role").Limit(pageSize).Offset(offset)
	var total int64
	ur.database.Table("users").Count(&total)
	err := db.Find(&persons).Error

	if err != nil {
		return nil, 0, fmt.Errorf("failed view all rooms %v", err.Error())
	}

	return &persons, total, nil
}

func (ur *userRepository) Delete(id uint) error {
	err := ur.database.Table("users").Where("id = ?", id).Delete(&domain.User{}).Error

	if err != nil {
		return err
	}

	return nil
}
