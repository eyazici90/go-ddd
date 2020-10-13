package order

import "context"

type Repository interface {
	GetOrders(context.Context) []*Order
	Get(ctx context.Context, id string) *Order
	Create(ctx context.Context, o *Order)
	Update(ctx context.Context, o *Order)
}
