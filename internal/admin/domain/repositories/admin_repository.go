//api_hexagonal_go/internal/admin/domain/repositories/admin_repository.go
package repositories

import "api-joaquin/internal/admin/domain"


type AdminRepository interface {
	Create(admin *domain.Admin) error
	GetAll() ([]domain.Admin, error)
	Update(admin *domain.Admin) error
	Delete(id int) error
}
