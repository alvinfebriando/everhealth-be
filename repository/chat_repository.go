package repository

import (
	"github.com/alvinfebriando/project-batman-be/entity"
	"gorm.io/gorm"
)

type ChatRepository interface {
	BaseRepository[entity.Chat]
}

type chatRepository struct {
	*baseRepository[entity.Chat]
	db *gorm.DB
}

func NewChatRepository(db *gorm.DB) ChatRepository {
	return &chatRepository{
		db:             db,
		baseRepository: &baseRepository[entity.Chat]{db: db},
	}
}
