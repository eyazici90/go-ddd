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

func (handler CancelOrderCommandHandler) Handle(_ context.Context, cmd CancelOrderCommand) error {
	order := handler.repository.Get(cmd.OrderId)
	order.Cancel()
	handler.repository.Update(order)
	return nil
}
