package query

type GetOrdersModel struct {
	OrderViews []OrderView
}

type GetOrderModel struct {
	OrderView OrderView
}

type OrderView struct {
	Id     string
	Status int
}
