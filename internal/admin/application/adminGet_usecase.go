//api_hexagonal_go/internal/admin/application/adminGet_usecase.go
package application

import (
	"api-joaquin/internal/admin/domain"
	"api-joaquin/internal/admin/domain/repositories"
)


type AdminGetUseCase struct {
	Repo repositories.AdminRepository
}


func NewAdminGetUseCase(repo repositories.AdminRepository) *AdminGetUseCase {
	return &AdminGetUseCase{Repo: repo}
}

func (uc *AdminGetUseCase) Execute() ([]domain.Admin, error) {
	return uc.Repo.GetAll()
}
