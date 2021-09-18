package command

import (
	"context"

	"ordercontext/internal/domain/order"

	"github.com/eyazici90/go-mediator/pkg/mediator"
)

type CancelOrder struct {
	OrderID string `validate:"required,min=10"`
}

func (CancelOrder) Key() string { return "CancelOrder" }

type CancelOrderHandler struct {
	orderHandler
}

func NewCancelOrderHandler(getOrder GetOrder, updateOrder UpdateOrder) CancelOrderHandler {
	return CancelOrderHandler{
		orderHandler: newOrderHandler(getOrder, updateOrder),
	}
}

func (h CancelOrderHandler) Handle(ctx context.Context, msg mediator.Message) error {
	cmd, ok := msg.(CancelOrder)
	if err := checkType(ok); err != nil {
		return err
	}
	return h.update(ctx, cmd.OrderID, func(o *order.Order) {
		o.Cancel()
	})
}
