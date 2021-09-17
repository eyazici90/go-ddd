package store

import (
	"context"
	"sync"

	"ordercontext/internal/domain/order"
)

type OrderInMemoryRepository struct {
	data  map[string]*order.Order
	mutex sync.RWMutex
}

func NewOrderInMemoryRepository() *OrderInMemoryRepository {
	return &OrderInMemoryRepository{
		data: make(map[string]*order.Order),
	}
}

func (i *OrderInMemoryRepository) GetAll(_ context.Context) ([]*order.Order, error) {
	i.mutex.RLock()
	defer i.mutex.RUnlock()

	var result []*order.Order

	for _, v := range i.data {
		result = append(result, v)
	}

	return result, nil
}

func (i *OrderInMemoryRepository) Get(_ context.Context, id string) (*order.Order, error) {
	i.mutex.RLock()
	defer i.mutex.RUnlock()

	return i.data[id], nil
}

func (i *OrderInMemoryRepository) Update(_ context.Context, o *order.Order) error {
	i.mutex.Lock()
	defer i.mutex.Unlock()

	i.data[o.ID()] = o
	return nil
}

func (i *OrderInMemoryRepository) Create(_ context.Context, o *order.Order) error {
	i.mutex.Lock()
	defer i.mutex.Unlock()

	i.data[o.ID()] = o
	return nil
}
