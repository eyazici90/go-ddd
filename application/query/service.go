package query

import (
	"context"
	"orderContext/domain/order"
)

type OrderQueryService interface {
	GetOrders(context.Context) GetOrdersDto

	GetOrder(ctx context.Context, id string) GetOrderDto
}

type service struct {
	repository order.Repository
}

func NewOrderQueryService(r order.Repository) OrderQueryService {
	return &service{r}
}

func (s *service) GetOrders(ctx context.Context) GetOrdersDto {
	oViews := mapToAll(s.repository.GetOrders(ctx))

	result := GetOrdersDto{oViews}

	return result
}

func (s *service) GetOrder(ctx context.Context, id string) GetOrderDto {
	oView := mapTo(s.repository.Get(ctx, id))

	result := GetOrderDto{oView}
	return result
}

func mapTo(o *order.Order) OrderView {
	return OrderView{o.Id(), int(o.Status())}
}

func mapToAll(orders []*order.Order) []OrderView {
	result := make([]OrderView, len(orders))

	for i, o := range orders {
		result[i] = mapTo(o)
	}
	return result
}
