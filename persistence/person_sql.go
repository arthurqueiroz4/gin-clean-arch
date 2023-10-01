package persistence

import (
	"echo-go/models"
	"fmt"

	"gorm.io/gorm"
)

type PersonSQL struct {
	DB *gorm.DB
}

func NewPersonSQL(DB *gorm.DB) *PersonSQL {
	return &PersonSQL{DB: DB}
}

func (prs *PersonSQL) Create(person *models.Person) (*models.Person, error){
	err := prs.DB.Create(&person).Error

	if err != nil {
		return nil, fmt.Errorf("failed create person %v", err.Error())
	}

	return person, nil
}

func (prs *PersonSQL) FindById(id int) (*models.Person, error) {
	var person = models.Person{}

	err := prs.DB.Where("id = ?", id).First(&person).Error

	if err != nil {
		return nil, fmt.Errorf("id not exists")
	}

	return &person, nil
}

func (prs *PersonSQL) FindByName(name string) (*models.Person, error){
	var person models.Person

	err := prs.DB.Where("name = ?", name).Find(&person).Error

	if err != nil {
		return nil, fmt.Errorf("name not exists")
	}

	return &person, nil
}

func (prs *PersonSQL) FindAll(page, pageSize int, nome string) (*[]models.Person, error) {
	var persons []models.Person

	offset := (page - 1) * pageSize

	db := prs.DB.Select("id", "nome", "email", "pass").Limit(pageSize).Offset(offset)

	err := db.Find(&persons).Error

	if err != nil {
		return nil, fmt.Errorf("failed view all rooms %v", err.Error())
	}

	return &persons, nil
}

func (prs *PersonSQL) Delete(id int) error {
	err := prs.DB.Delete(id).Error

	if err != nil {
		return fmt.Errorf("failed delete id:%d", id)
	}

	return nil
}

