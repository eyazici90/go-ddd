package query

import (
	"context"

	"ordercontext/internal/domain"
)

type OrderQueryService interface {
	GetOrders(context.Context) GetOrdersDto
	GetOrder(ctx context.Context, id string) GetOrderDto
}

type Service struct {
	repository domain.OrderRepository
}

func NewOrderQueryService(r domain.OrderRepository) *Service {
	return &Service{r}
}

func (s *Service) GetOrders(ctx context.Context) GetOrdersDto {
	orders, err := s.repository.GetAll(ctx)
	if err != nil {
		return GetOrdersDto{}
	}
	oViews := mapToAll(orders)

	result := GetOrdersDto{oViews}

	return result
}

func (s *Service) GetOrder(ctx context.Context, id string) GetOrderDto {
	order, err := s.repository.Get(ctx, id)
	if err != nil {
		return GetOrderDto{}
	}
	oView := mapTo(order)

	result := GetOrderDto{oView}
	return result
}
