package interfaces

import (
	"api-joaquin/internal/mailer/controllers"
	"api-joaquin/internal/mailer/domain"
	"github.com/gofiber/fiber/v2"
)

// MailHandler maneja las peticiones HTTP relacionadas con correos
type MailHandler struct {
	Controller *controllers.MailController
}

// NewMailHandler crea un nuevo manejador de correo
func NewMailHandler(controller *controllers.MailController) *MailHandler {
	return &MailHandler{Controller: controller}
}

// SendMail maneja el envío de correos simples
func (h *MailHandler) SendMail(c *fiber.Ctx) error {
	mail := new(domain.Mail)
	
	if err := c.BodyParser(mail); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Error al procesar la solicitud",
			"error":   err.Error(),
		})
	}

	// Validación simple (en producción usaría un validador como go-playground/validator)
	if mail.To == "" || mail.Subject == "" || mail.Body == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Todos los campos son requeridos: to, subject, body",
		})
	}

	if err := h.Controller.SimpleUseCase.Send(mail); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Error al enviar el correo",
			"error":   err.Error(),
		})
	}

	return c.JSON(domain.MailResponse{
		Success: true,
		Message: "Correo enviado exitosamente (simulado)",
	})
}

// SendMailWithTemplate maneja el envío de correos con plantilla
func (h *MailHandler) SendMailWithTemplate(c *fiber.Ctx) error {
	type request struct {
		To       string      `json:"to"`
		Subject  string      `json:"subject"`
		Template string      `json:"template"`
		Data     interface{} `json:"data"`
	}

	req := new(request)
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Error al procesar la solicitud",
			"error":   err.Error(),
		})
	}

	mail := &domain.Mail{
		To:      req.To,
		Subject: req.Subject,
	}

	if err := h.Controller.TemplateUseCase.SendWithTemplate(mail, req.Template, req.Data); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Error al enviar el correo con plantilla",
			"error":   err.Error(),
		})
	}

	return c.JSON(domain.MailResponse{
		Success: true,
		Message: "Correo con plantilla enviado exitosamente (simulado)",
	})
}