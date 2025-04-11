package infrastructure

import (
	"api-joaquin/internal/mailer/domain"
	"api-joaquin/internal/mailer/domain/repositories"
	"fmt"
	"log"
	"math/rand"
	"time"
)

type MailService struct {
	simulateDelay bool
}

func NewMailService() repositories.MailRepository {
	rand.Seed(time.Now().UnixNano())
	return &MailService{
		simulateDelay: true,
	}
}

func (s *MailService) Send(mail *domain.Mail) error {
	if s.simulateDelay {
		time.Sleep(500 * time.Millisecond)
	}

	if rand.Intn(5) == 0 {
		errMsg := fmt.Sprintf("error simulado en el servicio de correo (intento fallido para %s)", mail.To)
		log.Printf("⚠️ %s", errMsg)
		return fmt.Errorf(errMsg)
	}

	log.Printf(`
📧 === CORREO SIMULADO ===
📨 Para: %s
📌 Asunto: %s
📝 Mensaje:
%s
✅ Correo simulado enviado con éxito
=========================
`, mail.To, mail.Subject, mail.Body)

	return nil
}