package repositories

import "api-joaquin/internal/mailer/domain"

// MailRepository define el contrato para el envío de correos
type MailRepository interface {
	Send(mail *domain.Mail) error
	SendWithTemplate(mail *domain.Mail, template string, data interface{}) error
}