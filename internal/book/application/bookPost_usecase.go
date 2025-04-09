//api_hexagonal_go/internal/book/application/bookPost_usecase.go
package application

import (
	"api-joaquin/internal/book/domain"
	"api-joaquin/internal/book/domain/repositories"
)


type BookPostUseCase struct {
	Repo repositories.BookRepository
}


func NewBookPostUseCase(repo repositories.BookRepository) *BookPostUseCase {
	return &BookPostUseCase{Repo: repo}
}


func (uc *BookPostUseCase) Execute(book *domain.Book) error {
	return uc.Repo.Create(book)
}
