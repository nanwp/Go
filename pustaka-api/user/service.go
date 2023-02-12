package user

type Service interface {
	FindAll() ([]User, error)
	FindUser(username string) (User, error)
	// Create(user User) (User, error)
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
