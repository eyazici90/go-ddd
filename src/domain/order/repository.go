package order

import "context"

type OrderRepository interface {
	GetOrders(context.Context) []Order
	Get(ctx context.Context, id string) Order
	Create(ctx context.Context, o Order)
	Update(ctx context.Context, o Order)
}
