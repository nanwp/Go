package users

import "pustaka-api/middleware"

type Service interface {
	FindAll() ([]User, error)
	FindUser(username string) (User, error)
	Create(user UserRequest) (User, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindAll() ([]User, error) {
	return s.repository.FindAll()
}

func (s *service) FindUser(username string) (User, error) {
	return s.repository.FindUser(username)
}

func (s *service) Create(user UserRequest) (User, error) {

	pw := middleware.GetPwd(user.Password)
	hashPw := middleware.HashAndSalt(pw)

	users := User{
		Username: user.Username,
		Password: hashPw,
	}

	newUser, err := s.repository.Create(users)

	return newUser, err
}
