package command

import (
	"context"

	"ordercontext/internal/domain/order"

	"github.com/eyazici90/go-mediator"
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

func (handler PayOrderCommandHandler) Handle(ctx context.Context, request mediator.Message) error {
	cmd := request.(PayOrderCommand)
	return handler.update(ctx, cmd.OrderID, func(order *order.Order) {
		order.Pay()
	})
}
