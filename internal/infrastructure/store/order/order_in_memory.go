package order

import (
	"context"
	"sync"

	"ordercontext/internal/domain"
)

type InMemoryRepository struct {
	data  map[string]*domain.Order
	mutex sync.RWMutex
}

func NewInMemoryRepository() *InMemoryRepository {
	return &InMemoryRepository{
		data: make(map[string]*domain.Order),
	}
}

func (i *InMemoryRepository) GetAll(_ context.Context) ([]*domain.Order, error) {
	i.mutex.RLock()
	defer i.mutex.RUnlock()

	var result []*domain.Order

	for _, v := range i.data {
		result = append(result, v)
	}

	return result, nil
}

func (i *InMemoryRepository) Get(_ context.Context, id string) (*domain.Order, error) {
	i.mutex.RLock()
	defer i.mutex.RUnlock()

	return i.data[id], nil
}

func (i *InMemoryRepository) Update(_ context.Context, o *domain.Order) error {
	i.mutex.Lock()
	defer i.mutex.Unlock()

	i.data[string(o.ID())] = o
	return nil
}

func (i *InMemoryRepository) Create(_ context.Context, o *domain.Order) error {
	i.mutex.Lock()
	defer i.mutex.Unlock()

	i.data[string(o.ID())] = o
	return nil
}
