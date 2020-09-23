package command

import (
	"context"
	"orderContext/domain/order"
)

type CancelOrderCommand struct {
	OrderId string `validate:"required,min=10"`
}

type CancelOrderCommandHandler struct {
	repository order.OrderRepository
}

func (handler CancelOrderCommandHandler) Handle(ctx context.Context, cmd CancelOrderCommand) error {
	order := handler.repository.Get(ctx, cmd.OrderId)
	order.Cancel()
	handler.repository.Update(ctx, order)
	return nil
}
