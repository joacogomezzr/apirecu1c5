// api_hexagonal_go/internal/mailer/domain/mail.go
package domain

type Mail struct {
	To      string `json:"to"`
	Subject string `json:"subject"`
	Body    string `json:"body"`
}