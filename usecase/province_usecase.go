package usecase

import (
	"context"

	"github.com/alvinfebriando/project-batman-be/apperror"
	"github.com/alvinfebriando/project-batman-be/entity"
	"github.com/alvinfebriando/project-batman-be/repository"
)

type ProvinceUsecase interface {
	FindAllProvince(ctx context.Context) ([]*entity.Province, error)
	FindProvinceByIdWithCities(ctx context.Context, id uint) (*entity.Province, error)
}

type provinceUsecase struct {
	provinceRepository repository.ProvinceRepository
}

func NewProvinceUsecase(rp repository.ProvinceRepository) ProvinceUsecase {
	return &provinceUsecase{provinceRepository: rp}
}

func (u *provinceUsecase) FindAllProvince(ctx context.Context) ([]*entity.Province, error) {
	return u.provinceRepository.FindAllProvince(ctx)
}

func (u *provinceUsecase) FindProvinceByIdWithCities(ctx context.Context, id uint) (*entity.Province, error) {
	province, err := u.provinceRepository.FindProvinceDetail(ctx, id)
	if err != nil {
		return province, err
	}
	if province == nil {
		return province, apperror.NewResourceNotFoundError("province", "id", id)
	}
	return province, nil
}
