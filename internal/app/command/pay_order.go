// nolint:dupl
package command

import (
	"context"

	"github.com/eyazici90/go-ddd/internal/order"
	"github.com/eyazici90/go-mediator/mediator"
)

type PayOrder struct {
	OrderID string `validate:"required,min=10"`
}

func (PayOrder) Key() int { return payCommandKey }

type PayOrderHandler struct {
	orderHandler
}

func NewPayOrderHandler(getter OrderGetter, updater OrderUpdater) PayOrderHandler {
	return PayOrderHandler{
		orderHandler: newOrderHandler(getter, updater),
	}
}

func (h PayOrderHandler) Handle(ctx context.Context, msg mediator.Message) error {
	cmd, ok := msg.(PayOrder)
	if !ok {
		return ErrInvalidCommand
	}
	return h.update(ctx, cmd.OrderID, func(o *order.Order) {
		o.Pay()
	})
}
