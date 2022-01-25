// nolint:dupl
package command

import (
	"context"

	"github.com/eyazici90/go-ddd/internal/domain"

	"github.com/eyazici90/go-mediator/mediator"
)

type PayOrder struct {
	OrderID string `validate:"required,min=10"`
}

func (PayOrder) Key() int { return payCommandKey }

type PayOrderHandler struct {
	orderHandler
}

func NewPayOrderHandler(orderGetter OrderGetter, orderUpdater OrderUpdater) PayOrderHandler {
	return PayOrderHandler{
		orderHandler: newOrderHandler(orderGetter, orderUpdater),
	}
}

func (h PayOrderHandler) Handle(ctx context.Context, msg mediator.Message) error {
	cmd, ok := msg.(PayOrder)
	if err := checkType(ok); err != nil {
		return err
	}
	return h.update(ctx, cmd.OrderID, func(o *domain.Order) {
		o.Pay()
	})
}
