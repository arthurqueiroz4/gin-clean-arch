package domain

type JSONUser struct {
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required,email"`
	Pass  string `json:"password" binding:"required"`
	Role  string `json:"role" binding:"required"`
}

type User struct {
	ID    uint `gorm:"primaryKey"`
	Name  string
	Email string `gorm:"unique"`
	Pass  string
	Role  string
}

type UserRepository interface {
	Create(user *User) (*User, error)
	FindById(id uint) (*User, error)
	FindByEmail(email string) (*User, error)
	FindAll(page, pageSize int) (*[]User, error)
	Delete(id int) error
}

type UserUsecase interface {
	Create(user User) (*User, error)
	FindByEmail(email string) (*User, error)
	FindById(id uint) (*User, error)
	// FindByEmail(email string, userRepository UserRepository) (*User, error)
	// FindAll(page, pageSize int, userRepository UserRepository) (*[]User, error)
	// Delete(id int, userRepository UserRepository) error
}
