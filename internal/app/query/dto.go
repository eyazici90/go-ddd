package query

type (
	GetOrdersDto struct {
		OrderViews []OrderView
	}

	GetOrderDto struct {
		OrderView OrderView
	}

	OrderView struct {
		ID         string
		CustomerID string
		ProductID  string
		Status     int
	}
)
