package routes

import (
	"github.com/gofiber/fiber/v2"
	"api-joaquin/internal/mailer/interfaces"
)

// SetupMailRoutes configura las rutas para el servicio de correo
func SetupMailRoutes(app *fiber.App, handler *interfaces.MailHandler) {
	mailGroup := app.Group("/api/mail")
	
	// Rutas de correo
	mailGroup.Post("/send", handler.SendMail)
	mailGroup.Post("/send-template", handler.SendMailWithTemplate)
}