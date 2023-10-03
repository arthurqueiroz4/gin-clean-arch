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
	email, _ := ur.FindByEmail(user.Email)

	if email != nil {
		return nil, fmt.Errorf("email already registered")
	}

	err := ur.database.Create(&user).Error

	if err != nil {
		return nil, fmt.Errorf("failed create user %v", err.Error())
	}

	return user, nil
}

func (ur *userRepository) FindById(id int) (*domain.User, error) {
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

func (ur *userRepository) FindAll(page, pageSize int) (*[]domain.User, error) {
	var persons []domain.User

	offset := (page - 1) * pageSize

	db := ur.database.Select("id", "nome", "email", "pass").Limit(pageSize).Offset(offset)

	err := db.Find(&persons).Error

	if err != nil {
		return nil, fmt.Errorf("failed view all rooms %v", err.Error())
	}

	return &persons, nil
}

func (ur *userRepository) Delete(id int) error {
	err := ur.database.Delete(id).Error

	if err != nil {
		return fmt.Errorf("failed delete id:%d", id)
	}

	return nil
}
