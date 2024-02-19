package usecase

import (
	"context"

	"github.com/alvinfebriando/project-batman-be/entity"
	"github.com/alvinfebriando/project-batman-be/repository"
)

type DoctorSpecialistUsecase interface {
	FindAllDoctorSpecialist(ctx context.Context) ([]*entity.DoctorSpecialist, error)
}

type doctorSpecialistUsecase struct {
	doctorSpecialistRepo repository.DoctorSpecialistRepository
}

func NewDoctorSpecialistUsecase(r repository.DoctorSpecialistRepository) DoctorSpecialistUsecase {
	return &doctorSpecialistUsecase{doctorSpecialistRepo: r}
}

func (u *doctorSpecialistUsecase) FindAllDoctorSpecialist(ctx context.Context) ([]*entity.DoctorSpecialist, error) {
	return u.doctorSpecialistRepo.FindAllDoctorSpecialist(ctx)
}
