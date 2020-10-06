package order

type (
	OrderCreatedEvent struct {
		id string
	}
	OrderPaidEvent struct {
		id string
	}
	OrderCancelledEvent struct {
		id string
	}
)
