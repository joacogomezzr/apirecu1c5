package application

import (
	"api-joaquin/internal/mailer/domain"
	"api-joaquin/internal/mailer/domain/repositories"
)

// MailUseCase contiene la lógica de negocio para el envío de correos
type MailUseCase struct {
	Repo repositories.MailRepository
}

// NewMailUseCase crea una nueva instancia de MailUseCase
func NewMailUseCase(repo repositories.MailRepository) *MailUseCase {
	return &MailUseCase{Repo: repo}
}

// Send ejecuta el envío de un correo simple
func (uc *MailUseCase) Send(mail *domain.Mail) error {
	return uc.Repo.Send(mail)
}

// SendWithTemplate ejecuta el envío de un correo con plantilla
func (uc *MailUseCase) SendWithTemplate(mail *domain.Mail, template string, data interface{}) error {
	return uc.Repo.SendWithTemplate(mail, template, data)
}