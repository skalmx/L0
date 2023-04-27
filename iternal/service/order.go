package service

import (
	"L0/iternal/domain"
	"L0/pkg/cache"
	"context"
)

type OrderRepo interface {
	Create(ctx context.Context, orderUid string, order domain.Order) error
	GetById(ctx context.Context, orderUid string) (domain.Order, error)
}

type OrderService struct {
	repo OrderRepo
	cache *cache.MemoryCache
}

func NewOrderService(repo OrderRepo, cache *cache.MemoryCache) *OrderService {
	return &OrderService{
		repo: repo,
		cache: cache,
	}
}

func (o *OrderService) Create(ctx context.Context, orderUid string, order domain.Order) error {
	o.cache.Set(&order)
	return o.repo.Create(ctx, orderUid, order)
}

func (o *OrderService) GetById(ctx context.Context, orderUid string) (domain.Order, error) {
	return o.repo.GetById(ctx, orderUid)
}
