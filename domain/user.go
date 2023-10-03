package domain

type User struct {
	ID    string `gorm:"primarykey"`
	Name  string
	Email string `gorm:"unique"`
	Pass  string
}

type UserRepository interface {
	Create(user *User) (*User, error)
	FindById(id int) (*User, error)
	FindByEmail(email string) (*User, error)
	FindAll(page, pageSize int) (*[]User, error)
	Delete(id int) error
}
