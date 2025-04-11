// api_hexagonal_go/internal/book/application/bookPost_usecase.go
package application

import (
	"api-joaquin/internal/book/domain"
	"api-joaquin/internal/book/domain/repositories"
)

type BookPostUseCase struct {
	Repo  repositories.BookRepository
	Email repositories.IEmailRepository
}

func NewBookPostUseCase(repo repositories.BookRepository, email repositories.IEmailRepository) *BookPostUseCase {
	return &BookPostUseCase{
		Repo:  repo,
		Email: email,
	}
}

func (uc *BookPostUseCase) Execute(book *domain.Book) error {
	err := uc.Repo.Create(book)
	if err != nil {
		return err
	}
	err = uc.Email.SendMail(book)
	if err != nil {
		return err
	}

	return nil
}
