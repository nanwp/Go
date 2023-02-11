package entity

type Service interface {
	FindAll() ([]Book, error)
	FindByID(ID int) (Book, error)
	Create(book BookRequest) (Book, error)
}

type service struct {
	reppository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindALL() ([]Book, error) {
	books, err := s.reppository.FindAll()
	return books, err
	// return s.reppository.FindAll()
}

func (s *service) FindByID(ID int) (Book, error) {
	book, err := s.reppository.FindByID(ID)
	return book, err
	// return s.reppository.FindAll()
}

func (s *service) Create(bookRequest BookRequest) (Book, error) {
	price, _ := bookRequest.Price.Int64()

	book := Book{
		Title: bookRequest.Title,
		Price: int(price),
	}

	newbook, err := s.reppository.Create(book)
	return newbook, err
}
