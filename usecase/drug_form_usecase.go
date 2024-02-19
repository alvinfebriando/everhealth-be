package usecase

import (
	"context"

	"github.com/alvinfebriando/project-batman-be/entity"
	"github.com/alvinfebriando/project-batman-be/repository"
)

type DrugFormUsecase interface {
	FindAllDrugForm(ctx context.Context) ([]*entity.DrugForm, error)
}

type drugFormUsecase struct {
	drugFormRepo repository.DrugFormRepository
}

func NewDrugFormUsecase(r repository.DrugFormRepository) DrugFormUsecase {
	return &drugFormUsecase{drugFormRepo: r}
}

func (u *drugFormUsecase) FindAllDrugForm(ctx context.Context) ([]*entity.DrugForm, error) {
	return u.drugFormRepo.FindAllDrugForm(ctx)
}
