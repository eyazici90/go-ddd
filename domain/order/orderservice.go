package order

import (
	"strconv"
	"sync"
)

type OrderService interface {
	GetOrders() []Order
	Get(id string) Order
	Create(o Order)
}

var fakeOrders = []Order{}

var lockMutex = new(sync.RWMutex)

type service struct{}

func (s *service) GetOrders() []Order {
	return fakeOrders
}

func (s *service) Get(id string) Order {
	i, _ := strconv.Atoi(id)
	return fakeOrders[i]
}

func (s *service) Create(o Order) {
	lockMutex.Lock()
	defer lockMutex.Unlock()

	i, _ := strconv.Atoi(string(o.id))
	fakeOrders[i] = o
}
