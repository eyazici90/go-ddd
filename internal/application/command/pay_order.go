package command

import (
	"context"

	"ordercontext/internal/domain/order"

	"github.com/eyazici90/go-mediator/mediator"
)

type PayOrderCommand struct {
	OrderID string `validate:"required,min=10"`
}

func (PayOrderCommand) Key() string { return "PayOrderCommand" }

type PayOrderCommandHandler struct {
	commandHandlerBase
}

func NewPayOrderCommandHandler(getOrder GetOrder, updateOrder UpdateOrder) PayOrderCommandHandler {
	return PayOrderCommandHandler{
		commandHandlerBase: newcommandHandlerBase(getOrder, updateOrder),
	}
}

func (h PayOrderCommandHandler) Handle(ctx context.Context, req mediator.Message) error {
	cmd := req.(PayOrderCommand)
	return h.update(ctx, cmd.OrderID, func(o *order.Order) {
		o.Pay()
	})
}
