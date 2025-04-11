package domain

// Mail representa la estructura de un correo electrónico
type Mail struct {
	To      string `json:"to" validate:"required,email"`
	Subject string `json:"subject" validate:"required,min=3,max=100"`
	Body    string `json:"body" validate:"required,min=10"`
}

// MailResponse representa la respuesta del servicio de correo
type MailResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}