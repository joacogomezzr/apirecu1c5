//api_hexagonal_go/internal/book/domain/repositories/book_repository.go
package repositories

import "api-joaquin/internal/book/domain"


type BookRepository interface {
	Create(book *domain.Book) error
	GetAll() ([]domain.Book, error)
	Update(book *domain.Book) error
	Delete(id int) error
}
