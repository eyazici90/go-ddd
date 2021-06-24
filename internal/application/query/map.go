package query

import "ordercontext/internal/domain"

func mapTo(o *domain.Order) OrderView {
	return OrderView{ID: o.ID(), Status: int(o.Status()), ProductID: o.ProductID(), CustomerID: o.CustomerID()}
}

func mapToAll(orders []*domain.Order) []OrderView {
	result := make([]OrderView, len(orders))

	for i, o := range orders {
		result[i] = mapTo(o)
	}
	return result
}
