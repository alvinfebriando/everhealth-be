package repository

import (
	"github.com/alvinfebriando/project-batman-be/entity"
	"gorm.io/gorm"
)

type RoleRepository interface {
	BaseRepository[entity.Role]
}

type roleRepository struct {
	*baseRepository[entity.Role]
	db *gorm.DB
}

func NewRoleRepository(db *gorm.DB) RoleRepository {
	return &roleRepository{
		db:             db,
		baseRepository: &baseRepository[entity.Role]{db: db},
	}
}
