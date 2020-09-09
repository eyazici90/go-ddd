package order

type OrderRepository interface {
	GetOrders() []Order
	Get(id string) Order
	Create(o Order)
	Update(o Order)
}
