package domain

import "context"

type OrderRepository interface {
	GetAll(context.Context) ([]*Order, error)
	Get(ctx context.Context, id string) (*Order, error)
	Create(ctx context.Context, o *Order) error
	Update(ctx context.Context, o *Order) error
}
