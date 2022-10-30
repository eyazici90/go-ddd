package inmem

import (
	"context"
	"sync"

	"github.com/eyazici90/go-ddd/internal/order"
)

type OrderRepository struct {
	data  map[string]*order.Order
	mutex sync.RWMutex
}

func NewOrderRepository() *OrderRepository {
	return &OrderRepository{
		data: make(map[string]*order.Order),
	}
}

func (i *OrderRepository) GetAll(_ context.Context) ([]*order.Order, error) {
	i.mutex.RLock()
	defer i.mutex.RUnlock()

	var result []*order.Order

	for _, v := range i.data {
		result = append(result, v)
	}

	return result, nil
}

func (i *OrderRepository) Get(_ context.Context, id string) (*order.Order, error) {
	i.mutex.RLock()
	defer i.mutex.RUnlock()

	return i.data[id], nil
}

func (i *OrderRepository) Update(_ context.Context, o *order.Order) error {
	i.mutex.Lock()
	defer i.mutex.Unlock()

	i.data[o.ID()] = o
	return nil
}

func (i *OrderRepository) Create(_ context.Context, o *order.Order) error {
	i.mutex.Lock()
	defer i.mutex.Unlock()

	i.data[o.ID()] = o
	return nil
}
