package usecase

import (
	"context"

	"github.com/alvinfebriando/project-batman-be/entity"
	"github.com/alvinfebriando/project-batman-be/repository"
)

type OrderStatusUsecase interface {
	FindAllOrderStatus(ctx context.Context) ([]*entity.OrderStatus, error)
}

type orderStatusUsecase struct {
	orderStatusRepository repository.OrderStatusRepository
}

func NewOrderStatusUsecase(rp repository.OrderStatusRepository) OrderStatusUsecase {
	return &orderStatusUsecase{orderStatusRepository: rp}
}

func (u *orderStatusUsecase) FindAllOrderStatus(ctx context.Context) ([]*entity.OrderStatus, error) {
	return u.orderStatusRepository.FindAllOrderStatus(ctx)
}
