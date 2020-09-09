package infrastructure

import (
	"orderContext/domain/order"
	"sync"
)

var fakeOrders = make(map[string]order.Order)

var lockMutex = new(sync.RWMutex)

type repository struct{}

func NewOrderRepository() order.OrderRepository {
	return &repository{}
}

func (r *repository) GetOrders() []order.Order {
	result := make([]order.Order, 0, len(fakeOrders))

	for _, v := range fakeOrders {
		result = append(result, v)
	}

	return result
}

func (r *repository) Get(id string) order.Order {
	return fakeOrders[id]
}

func (r *repository) Update(o order.Order) {
	lockMutex.Lock()
	defer lockMutex.Unlock()

	fakeOrders[string(o.ID)] = o
}

func (r *repository) Create(o order.Order) {
	lockMutex.Lock()
	defer lockMutex.Unlock()

	fakeOrders[string(o.ID)] = o
}
