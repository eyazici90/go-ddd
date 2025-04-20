package query

import (
	"github.com/eyazici90/go-ddd/internal/order"
)

func mapTo(o *order.Order) OrderView {
	return OrderView{ID: o.ID(), Status: int(o.Status()), ProductID: o.ProductID(), CustomerID: o.CustomerID()}
}

func mapToAll(orders []*order.Order) []OrderView {
	result := make([]OrderView, len(orders))

	for i, o := range orders {
		result[i] = mapTo(o)
	}
	return result
}
