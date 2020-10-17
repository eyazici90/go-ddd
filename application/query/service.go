package query

import (
	"context"
	"orderContext/domain/order"
)

type OrderQueryService interface {
	GetOrders(context.Context) GetOrdersDto

	GetOrder(ctx context.Context, id string) GetOrderDto
}

type Service struct {
	repository order.Repository
}

func NewOrderQueryService(r order.Repository) *Service {
	return &Service{r}
}

func (s *Service) GetOrders(ctx context.Context) GetOrdersDto {
	oViews := mapToAll(s.repository.GetOrders(ctx))

	result := GetOrdersDto{oViews}

	return result
}

func (s *Service) GetOrder(ctx context.Context, id string) GetOrderDto {
	oView := mapTo(s.repository.Get(ctx, id))

	result := GetOrderDto{oView}
	return result
}
