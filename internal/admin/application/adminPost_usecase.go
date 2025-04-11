//api_hexagonal_go/internal/admin/application/adminPost_usecase.go
package application

import (
	adminDomain "api-joaquin/internal/admin/domain"
	"api-joaquin/internal/admin/domain/repositories"
	mailerDomain "api-joaquin/internal/mailer/domain"
	mailRepo "api-joaquin/internal/mailer/domain/repositories"
	"fmt"
	"log"
	"time"
)

type AdminPostUseCase struct {
	AdminRepo repositories.AdminRepository
	MailRepo  mailRepo.MailRepository
}

func NewAdminPostUseCase(adminRepo repositories.AdminRepository, mailRepo mailRepo.MailRepository) *AdminPostUseCase {
	return &AdminPostUseCase{
		AdminRepo: adminRepo,
		MailRepo:  mailRepo,
	}
}

func (uc *AdminPostUseCase) Execute(admin *adminDomain.Admin) error {
	log.Printf("Intentando registrar nuevo admin: %s <%s>", admin.Name, admin.Email)

	if err := uc.AdminRepo.Create(admin); err != nil {
		log.Printf("❌ Error al crear admin: %v", err)
		return fmt.Errorf("error al crear administrador: %w", err)
	}

	mail := &mailerDomain.Mail{
		To:      admin.Email,
		Subject: "✅ Registro exitoso - Sistema de Biblioteca",
		Body: fmt.Sprintf(`
Estimado/a %s,

Su registro como administrador en el Sistema de Biblioteca se ha completado exitosamente.

Detalles de su cuenta:
• Nombre: %s
• Email: %s
• Fecha: %s

Atentamente,
El equipo de la Biblioteca
`, admin.Name, admin.Name, admin.Email, time.Now().Format("02/01/2006")),
	}

	if err := uc.MailRepo.Send(mail); err != nil {
		log.Printf("⚠️ Correo NO enviado a %s: %v", admin.Email, err)
	} else {
		log.Printf("✉️ Correo de confirmación enviado a %s", admin.Email)
	}

	return nil
}