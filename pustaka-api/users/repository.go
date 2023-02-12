package users

import "gorm.io/gorm"

type Repository interface {
	FindAll() ([]User, error)
	FindUser(username string) (User, error)
	Create(user User) (User, error)
	// Create(user User) (User, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]User, error) {
	var users []User
	err := r.db.Table("tbl_user").Find(&users).Error
	return users, err

}

func (r *repository) FindUser(username string) (User, error) {
	var user User
	err := r.db.Table("tbl_user").Find(&user, username).Error
	return user, err
}

func (r *repository) Create(user User) (User, error) {
	err := r.db.Table("tbl_user").Create(&user).Error
	return user, err
}
