package infrastructure

import (
	"api-joaquin/internal/mailer/domain"
	"api-joaquin/internal/mailer/domain/repositories"
	"fmt"
	"log"
	"time"
)

// MailService implementa el repositorio de correo con simulación
type MailService struct {
	simulateDelay bool // Para simular retraso en el envío
}

// NewMailService crea una nueva instancia del servicio de correo
func NewMailService() repositories.MailRepository {
	return &MailService{
		simulateDelay: true,
	}
}

// Send implementa el envío de correo simulado
func (s *MailService) Send(mail *domain.Mail) error {
	if s.simulateDelay {
		time.Sleep(1 * time.Second) // Simula retraso en el envío
	}

	log.Printf("✉️ [MAIL SERVICE] Simulando envío de correo a: %s\n", mail.To)
	log.Printf("📌 Asunto: %s\n", mail.Subject)
	log.Printf("📝 Contenido: %s\n", mail.Body)
	log.Println("✅ Correo simulado enviado con éxito")

	return nil
}

// SendWithTemplate implementa el envío de correo con plantilla simulado
func (s *MailService) SendWithTemplate(mail *domain.Mail, template string, data interface{}) error {
	if s.simulateDelay {
		time.Sleep(1 * time.Second)
	}

	log.Printf("✉️ [MAIL SERVICE] Simulando envío de correo con plantilla a: %s\n", mail.To)
	log.Printf("📌 Asunto: %s\n", mail.Subject)
	log.Printf("📝 Plantilla usada: %s\n", template)
	log.Printf("📊 Datos de la plantilla: %+v\n", data)
	log.Println("✅ Correo con plantilla simulado enviado con éxito")

	return nil
}

// simulateExternalService simula una llamada a un servicio externo
func (s *MailService) simulateExternalService() error {
	// Simulación de posibles errores (10% de probabilidad de error)
	if time.Now().Unix()%10 == 0 {
		return fmt.Errorf("error simulado en el servicio de correo externo")
	}
	return nil
}