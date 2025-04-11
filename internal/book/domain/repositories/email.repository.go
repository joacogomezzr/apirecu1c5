package repositories

import "api-joaquin/internal/book/domain"

type IEmailRepository interface {
	SendMail(book *domain.Book) error
}