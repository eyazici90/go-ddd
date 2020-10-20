package query

import "orderContext/domain/order"

func mapTo(o *order.Order) OrderView {
	return OrderView{Id: o.ID(), Status: int(o.Status()), ProductId: o.ProductId(), CustomerId: o.CustomerId()}
}

func mapToAll(orders []*order.Order) []OrderView {
	result := make([]OrderView, len(orders))

	for i, o := range orders {
		result[i] = mapTo(o)
	}
	return result
}
