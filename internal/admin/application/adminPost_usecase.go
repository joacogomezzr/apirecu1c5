//api_hexagonal_go/internal/admin/application/adminPost_usecase.go
package application

import (
	"api-joaquin/internal/admin/domain"
	"api-joaquin/internal/admin/domain/repositories"
)
type AdminPostUseCase struct {
	Repo repositories.AdminRepository
}


func NewAdminPostUseCase(repo repositories.AdminRepository) *AdminPostUseCase {
	return &AdminPostUseCase{Repo: repo}
}


func (uc *AdminPostUseCase) Execute(admin *domain.Admin) error {
	return uc.Repo.Create(admin)
}
