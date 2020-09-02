package order

import (
	"sync"
)

type OrderRepository interface {
	GetOrders() []Order
	Get(id string) Order
	Create(o Order)
	Update(o Order)
}

var fakeOrders = make(map[string]Order)

var lockMutex = new(sync.RWMutex)

type repository struct{}

func NewOrderRepository() OrderRepository {
	return &repository{}
}

func (r *repository) GetOrders() []Order {
	result := make([]Order, 0, len(fakeOrders))

	for _, v := range fakeOrders {
		result = append(result, v)
	}

	return result
}

func (r *repository) Get(id string) Order {
	return fakeOrders[id]
}

func (r *repository) Update(o Order) {
	lockMutex.Lock()
	defer lockMutex.Unlock()

	fakeOrders[string(o.ID)] = o
}

func (r *repository) Create(o Order) {
	lockMutex.Lock()
	defer lockMutex.Unlock()

	fakeOrders[string(o.ID)] = o
}
