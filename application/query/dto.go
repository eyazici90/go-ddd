package query

type (
	GetOrdersDto struct {
		OrderViews []OrderView
	}

	GetOrderDto struct {
		OrderView OrderView
	}

	OrderView struct {
		Id     string
		Status int
	}
)
