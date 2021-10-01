package command

import (
	"context"

	"ordercontext/internal/domain/order"

	"github.com/eyazici90/go-mediator/pkg/mediator"
)

type PayOrder struct {
	OrderID string `validate:"required,min=10"`
}

func (PayOrder) Key() string { return "PayOrder" }

type PayOrderHandler struct {
	orderHandler
}

func NewPayOrderHandler(getOrder GetOrderFunc, updateOrder UpdateOrderFunc) PayOrderHandler {
	return PayOrderHandler{
		orderHandler: newOrderHandler(getOrder, updateOrder),
	}
}

func (h PayOrderHandler) Handle(ctx context.Context, msg mediator.Message) error {
	cmd, ok := msg.(PayOrder)
	if err := checkType(ok); err != nil {
		return err
	}
	return h.update(ctx, cmd.OrderID, func(o *order.Order) {
		o.Pay()
	})
}
