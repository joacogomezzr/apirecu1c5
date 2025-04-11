package infrastructure

import (
	"api-joaquin/internal/book/domain"
	"api-joaquin/internal/book/domain/repositories"
	"fmt"
	"log"
	"net/smtp"
)

type EmailRepository struct {
	smtpHost     string
	smtpPort     string
	smtpUsername string
	smtpPassword string
	from         string
	to           string
}

func NewEmailRepository() repositories.IEmailRepository {
	return &EmailRepository{
		smtpHost:     "smtp.gmail.com",        // Servidor SMTP de Gmail
		smtpPort:     "587",                   // Puerto estándar para TLS
		smtpUsername: "joaquin@gmail.com",    // Tu dirección de correo
		smtpPassword: "joaquinsb",       // Tu contraseña o token de aplicación
		from:         "biblioteca@gmail.com",  // Dirección de correo del remitente
		to:           "admin@biblioteca.com",  // Destinatario predeterminado
	}
}

func (e *EmailRepository) SendMail(book *domain.Book) error {
	addr := fmt.Sprintf("%s:%s", e.smtpHost, e.smtpPort)
	auth := smtp.PlainAuth("", e.smtpUsername, e.smtpPassword, e.smtpHost)

	// Construir el asunto y cuerpo del correo
	subject := fmt.Sprintf("Nuevo libro registrado: %s", book.Title)
	body := fmt.Sprintf(`
Estimado administrador,

Se ha registrado un nuevo libro en el sistema:

ID: %d
Título: %s
Autor: %s
Año: %d

Este libro ha sido agregado a la base de datos de la biblioteca.

Saludos cordiales,
Sistema de Biblioteca
`, book.ID, book.Title, book.Author, book.Year)

	// Formatear el mensaje de correo electrónico
	message := []byte(fmt.Sprintf("To: %s\r\n"+
		"From: %s\r\n"+
		"Subject: %s\r\n"+
		"MIME-Version: 1.0\r\n"+
		"Content-Type: text/plain; charset=UTF-8\r\n"+
		"\r\n"+
		"%s\r\n", e.to, e.from, subject, body))

	// Enviar el correo electrónico
	err := smtp.SendMail(addr, auth, e.from, []string{e.to}, message)
	if err != nil {
		log.Printf("Error al enviar el correo electrónico: %v", err)
		return err
	}

	log.Printf("Correo electrónico enviado exitosamente a %s sobre el libro: %s", e.to, book.Title)
	return nil
}