package book

type Service interface {
	FindAll() ([]Book, error)
	FindByID(ID int) (Book, error)
	Create(book BookRequest) (Book, error)
	Update(ID int, bookRequestUpdate BookRequestUpdate) (Book, error)
	Delete(ID int) (Book, error)
}

type service struct {
	reppository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindAll() ([]Book, error) {
	return s.reppository.FindAll()
}

func (s *service) FindByID(ID int) (Book, error) {
	return s.reppository.FindByID(ID)
}

func (s *service) Create(bookRequest BookRequest) (Book, error) {
	price, _ := bookRequest.Price.Int64()
	rating, _ := bookRequest.Rating.Int64()

	book := Book{
		Title:       bookRequest.Title,
		Price:       int(price),
		Description: bookRequest.Description,
		Rating:      int(rating),
	}

	newBook, err := s.reppository.Create(book)
	return newBook, err
}

func (s *service) Update(ID int, bookRequestUpdate BookRequestUpdate) (Book, error) {
	book, _ := s.reppository.FindByID(ID)

	price, _ := bookRequestUpdate.Price.Int64()
	rating, _ := bookRequestUpdate.Rating.Int64()

	if bookRequestUpdate.Title != "" {
		book.Title = bookRequestUpdate.Title
	}
	if bookRequestUpdate.Description != "" {
		book.Description = bookRequestUpdate.Description
	}
	if bookRequestUpdate.Price != "" {
		book.Price = int(price)
	}
	if bookRequestUpdate.Rating != "" {
		book.Rating = int(rating)
	}
	updateBook, err := s.reppository.Update(book)
	return updateBook, err

}

func (s *service) Delete(ID int) (Book, error) {
	book, _ := s.reppository.FindByID(ID)

	deleteBook, err := s.reppository.Delete(book)
	return deleteBook, err

}
