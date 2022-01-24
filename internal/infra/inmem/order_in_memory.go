package inmem

import (
	"context"
	"sync"

	"github.com/eyazici90/go-ddd/internal/domain"
)

type OrderRepository struct {
	data  map[string]*domain.Order
	mutex sync.RWMutex
}

func NewOrderRepository() *OrderRepository {
	return &OrderRepository{
		data: make(map[string]*domain.Order),
	}
}

func (i *OrderRepository) GetAll(_ context.Context) ([]*domain.Order, error) {
	i.mutex.RLock()
	defer i.mutex.RUnlock()

	var result []*domain.Order

	for _, v := range i.data {
		result = append(result, v)
	}

	return result, nil
}

func (i *OrderRepository) Get(_ context.Context, id string) (*domain.Order, error) {
	i.mutex.RLock()
	defer i.mutex.RUnlock()

	return i.data[id], nil
}

func (i *OrderRepository) Update(_ context.Context, o *domain.Order) error {
	i.mutex.Lock()
	defer i.mutex.Unlock()

	i.data[o.ID()] = o
	return nil
}

func (i *OrderRepository) Create(_ context.Context, o *domain.Order) error {
	i.mutex.Lock()
	defer i.mutex.Unlock()

	i.data[o.ID()] = o
	return nil
}
