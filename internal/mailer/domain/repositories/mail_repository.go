// internal/mailer/domain/repositories/mail_repository.go
package repositories

import "api-joaquin/internal/mailer/domain"

type MailRepository interface {
	Send(mail *domain.Mail) error
}