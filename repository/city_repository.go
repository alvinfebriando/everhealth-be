package repository

import (
	"github.com/alvinfebriando/project-batman-be/entity"
	"gorm.io/gorm"
)

type CityRepository interface {
	BaseRepository[entity.City]
}

type cityRepository struct {
	*baseRepository[entity.City]
	db *gorm.DB
}

func NewCityRepository(db *gorm.DB) CityRepository {
	return &cityRepository{
		db:             db,
		baseRepository: &baseRepository[entity.City]{db: db},
	}
}
