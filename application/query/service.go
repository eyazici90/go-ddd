package query

import (
	"context"
	"orderContext/domain/order"
)

type OrderQueryService interface {
	GetOrders(context.Context) GetOrdersModel

	GetOrder(ctx context.Context, id string) GetOrderModel
}

type service struct {
	repository order.OrderRepository
}

func NewOrderQueryService(r order.OrderRepository) OrderQueryService {
	return &service{repository: r}
}

func (s *service) GetOrders(ctx context.Context) GetOrdersModel {
	oViews := mapToAll(s.repository.GetOrders(ctx))

	result := GetOrdersModel{OrderViews: oViews}

	return result
}

func (s *service) GetOrder(ctx context.Context, id string) GetOrderModel {
	oView := mapTo(s.repository.Get(ctx, id))

	result := GetOrderModel{OrderView: oView}
	return result
}

func mapTo(o order.Order) OrderView {
	return OrderView{
		Id:     o.ID,
		Status: int(o.Status()),
	}
}

func mapToAll(orders []order.Order) []OrderView {
	result := make([]OrderView, len(orders))

	for i, o := range orders {
		result[i] = mapTo(o)
	}
	return result
}
