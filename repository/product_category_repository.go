package repository

import (
	"context"

	"github.com/alvinfebriando/project-batman-be/entity"
	"github.com/alvinfebriando/project-batman-be/valueobject"
	"gorm.io/gorm"
)

type ProductCategoryRepository interface {
	BaseRepository[entity.ProductCategory]
	FindProductCategories(ctx context.Context, query *valueobject.Query) (*valueobject.PagedResult, error)
}

type productCategoryRepository struct {
	*baseRepository[entity.ProductCategory]
	db *gorm.DB
}

func NewProductCategoryRepository(db *gorm.DB) ProductCategoryRepository {
	return &productCategoryRepository{
		db:             db,
		baseRepository: &baseRepository[entity.ProductCategory]{db: db},
	}
}

func (r *productCategoryRepository) FindProductCategories(ctx context.Context, query *valueobject.Query) (*valueobject.PagedResult, error) {
	return r.paginate(ctx, query, func(db *gorm.DB) *gorm.DB {
		name := query.GetConditionValue("name")

		if name != nil {
			db.Where("name ILIKE ?", name)
		}
		isDrug := query.GetConditionValue("is_drug")
		if isDrug != nil {
			db.Where("is_drug = ?", isDrug)
		}

		return db
	})
}
