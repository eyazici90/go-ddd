package order

import (
	"context"
	"sync"

	"ordercontext/internal/domain/order"
)

var fakeOrders = make(map[string]*order.Order)

var lockMutex = &sync.RWMutex{}

type repository struct{}

var InMemoryRepository order.Repository = &repository{}

func (r *repository) GetOrders(_ context.Context) ([]*order.Order, error) {
	lockMutex.RLock()
	defer lockMutex.RUnlock()

	var result []*order.Order

	for _, v := range fakeOrders {
		result = append(result, v)
	}

	return result, nil
}

func (r *repository) Get(_ context.Context, id string) (*order.Order, error) {
	lockMutex.RLock()
	defer lockMutex.RUnlock()

	return fakeOrders[id], nil
}

func (r *repository) Update(_ context.Context, o *order.Order) error {
	lockMutex.Lock()
	defer lockMutex.Unlock()

	fakeOrders[string(o.ID())] = o
	return nil
}

func (r *repository) Create(_ context.Context, o *order.Order) error {
	lockMutex.Lock()
	defer lockMutex.Unlock()

	fakeOrders[string(o.ID())] = o
	return nil
}
