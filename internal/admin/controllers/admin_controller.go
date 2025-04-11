//api_hexagonal_go/internal/admin/controllers/admin_controller.go
package controllers

import (
	"api-joaquin/internal/admin/application"
	"api-joaquin/internal/admin/domain/repositories"
	mailRepo "api-joaquin/internal/mailer/domain/repositories"
)

type AdminController struct {
	PostUseCase   *application.AdminPostUseCase
	GetUseCase    *application.AdminGetUseCase
	PutUseCase    *application.AdminPutUseCase
	DeleteUseCase *application.AdminDeleteUseCase
}

func NewAdminController(repo repositories.AdminRepository, mailRepo mailRepo.MailRepository) *AdminController {
	return &AdminController{
		PostUseCase:   application.NewAdminPostUseCase(repo, mailRepo),
		GetUseCase:    application.NewAdminGetUseCase(repo),
		PutUseCase:    application.NewAdminPutUseCase(repo),
		DeleteUseCase: application.NewAdminDeleteUseCase(repo),
	}
}