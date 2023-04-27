package cache

import (
	"L0/iternal/domain"
	"errors"
	"sync"
)

type MemoryCache struct {
	sync.RWMutex
	Cache map[string]*domain.Order
}

func NewCache() *MemoryCache {
	c := &MemoryCache{}
	c.Cache = make(map[string]*domain.Order)
	
	return c
}

func (c *MemoryCache) Set(order *domain.Order) error {
	c.Lock()
	defer c.Unlock()

	c.Cache[order.OrderUID] = order

	return nil
}

func (c *MemoryCache) Get(key string) (*domain.Order, error) {
	c.RLock()
	data, ex := c.Cache[key]
	defer c.RUnlock()

	if !ex {
		return &domain.Order{}, errors.New("no such element in mememory cache")
	}

	return data, nil
}

