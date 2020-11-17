package query

import "ordercontext/internal/domain/order"

func mapTo(o *order.Order) OrderView {
	return OrderView{Id: o.ID(), Status: int(o.Status()), ProductId: o.ProductID(), CustomerId: o.CustomerID()}
}

func mapToAll(orders []*order.Order) []OrderView {
	result := make([]OrderView, len(orders))

	for i, o := range orders {
		result[i] = mapTo(o)
	}
	return result
}
