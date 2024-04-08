// nolint:dupl
package command

import (
	"context"

	"github.com/eyazici90/go-ddd/internal/domain"

	"github.com/eyazici90/go-mediator/mediator"
)

type CancelOrder struct {
	OrderID string `validate:"required,min=10"`
}

func (CancelOrder) Key() int { return cancelCommandKey }

type CancelOrderHandler struct {
	orderHandler
}

func NewCancelOrderHandler(getter OrderGetter, updater OrderUpdater) CancelOrderHandler {
	return CancelOrderHandler{
		orderHandler: newOrderHandler(getter, updater),
	}
}

func (h CancelOrderHandler) Handle(ctx context.Context, msg mediator.Message) error {
	cmd, ok := msg.(CancelOrder)
	if !ok {
		return ErrInvalidCommand
	}
	return h.update(ctx, cmd.OrderID, func(o *domain.Order) {
		o.Cancel()
	})
}
