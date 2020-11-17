package order

import "context"

type Repository interface {
	GetOrders(context.Context) ([]*Order, error)
	Get(ctx context.Context, id string) (*Order, error)
	Create(ctx context.Context, o *Order) error
	Update(ctx context.Context, o *Order) error
}
