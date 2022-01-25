// nolint:dupl
package command

import (
	"context"

	"github.com/eyazici90/go-ddd/internal/domain"

	"github.com/eyazici90/go-mediator/pkg/mediator"
)

type CancelOrder struct {
	OrderID string `validate:"required,min=10"`
}

func (CancelOrder) Key() string { return "CancelOrder" }

type CancelOrderHandler struct {
	orderHandler
}

func NewCancelOrderHandler(orderGetter OrderGetter, orderUpdater OrderUpdater) CancelOrderHandler {
	return CancelOrderHandler{
		orderHandler: newOrderHandler(orderGetter, orderUpdater),
	}
}

func (h CancelOrderHandler) Handle(ctx context.Context, msg mediator.Message) error {
	cmd, ok := msg.(CancelOrder)
	if err := checkType(ok); err != nil {
		return err
	}
	return h.update(ctx, cmd.OrderID, func(o *domain.Order) {
		o.Cancel()
	})
}
