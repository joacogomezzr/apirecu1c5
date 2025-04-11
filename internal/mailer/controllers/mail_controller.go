package controllers

import (
	"api-joaquin/internal/mailer/application"
	"api-joaquin/internal/mailer/domain/repositories"
)

// MailController maneja las operaciones de correo
type MailController struct {
	SimpleUseCase    *application.MailUseCase
	TemplateUseCase  *application.MailUseCase
}

// NewMailController crea un nuevo controlador de correo
func NewMailController(repo repositories.MailRepository) *MailController {
	return &MailController{
		SimpleUseCase:    application.NewMailUseCase(repo),
		TemplateUseCase:  application.NewMailUseCase(repo),
	}
}