package command

import (
	"context"

	"ordercontext/internal/domain"

	"github.com/eyazici90/go-mediator/mediator"
)

type PayOrderCommand struct {
	OrderID string `validate:"required,min=10"`
}

func (PayOrderCommand) Key() string { return "PayOrderCommand" }

type PayOrderCommandHandler struct {
	commandHandler
}

func NewPayOrderCommandHandler(getOrder GetOrder, updateOrder UpdateOrder) PayOrderCommandHandler {
	return PayOrderCommandHandler{
		commandHandler: newcommandHandlerBase(getOrder, updateOrder),
	}
}

func (h PayOrderCommandHandler) Handle(ctx context.Context, msg mediator.Message) error {
	cmd, ok := msg.(PayOrderCommand)
	if err := checkType(ok); err != nil {
		return err
	}
	return h.update(ctx, cmd.OrderID, func(o *domain.Order) {
		o.Pay()
	})
}
