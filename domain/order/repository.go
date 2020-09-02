package order

import (
	"strconv"
	"sync"
)

type OrderRepository interface {
	GetOrders() []Order
	Get(id string) Order
	Create(o Order)
	Update(o Order)
}

var fakeOrders = []Order{}

var lockMutex = new(sync.RWMutex)

type repository struct{}

func NewOrderRepository() OrderRepository {
	return &repository{}
}

func (r *repository) GetOrders() []Order {
	return fakeOrders
}

func (r *repository) Get(id string) Order {
	i, _ := strconv.Atoi(id)
	return fakeOrders[i]
}

func (r *repository) Update(o Order) {
	lockMutex.Lock()
	defer lockMutex.Unlock()

	i, _ := strconv.Atoi(string(o.ID))
	fakeOrders[i] = o
}

func (r *repository) Create(o Order) {
	lockMutex.Lock()
	defer lockMutex.Unlock()

	i, _ := strconv.Atoi(string(o.ID))
	fakeOrders[i] = o
}
