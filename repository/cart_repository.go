package repository

import (
	"github.com/alvinfebriando/project-batman-be/entity"
	"gorm.io/gorm"
)

type CartRepository interface {
	BaseRepository[entity.Cart]
}

type cartRepository struct {
	*baseRepository[entity.Cart]
	db *gorm.DB
}

func NewCartRepository(db *gorm.DB) CartRepository {
	return &cartRepository{
		db:             db,
		baseRepository: &baseRepository[entity.Cart]{db: db},
	}
}
