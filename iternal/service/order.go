package service

import (
	"L0/iternal/domain"
	//"L0/iternal/repository"
	"context"
)

type OrderRepo interface {
	Create(ctx context.Context, orderUid string, order domain.Order) error
}

type OrderService struct {
	repo OrderRepo
}

func NewOrderService(repo OrderRepo) *OrderService {
	return &OrderService{
		repo: repo,
	}
}

func (o *OrderService) Create(ctx context.Context, orderUid string, order domain.Order) error {
	return o.repo.Create(ctx, orderUid, order)
}