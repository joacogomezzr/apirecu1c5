//api_hexagonal_go/internal/book/controllers/book_controller.go
package controllers

import (
	"api-joaquin/internal/book/application"
	"api-joaquin/internal/book/domain/repositories"
)


type BookController struct {
	PostUseCase   *application.BookPostUseCase
	GetUseCase    *application.BookGetUseCase
	PutUseCase    *application.BookPutUseCase
	DeleteUseCase *application.BookDeleteUseCase
}

func NewBookController(repo repositories.BookRepository, email repositories.IEmailRepository) *BookController {
	return &BookController{
		PostUseCase:   application.NewBookPostUseCase(repo, email),
		GetUseCase:    application.NewBookGetUseCase(repo),
		PutUseCase:    application.NewBookPutUseCase(repo),
		DeleteUseCase: application.NewBookDeleteUseCase(repo),
	}
}
