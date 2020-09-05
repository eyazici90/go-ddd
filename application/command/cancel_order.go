package command

import (
	"errors"
	"orderContext/domain/order"
)

type CancelOrderCommand struct {
	OrderId string
}

type CancelOrderCommandHandler struct {
	repository order.OrderRepository
}

type CancelOrderCommandValidator struct {
	cmd CancelOrderCommand
}

func (handler CancelOrderCommandHandler) Handle(cmd CancelOrderCommand) {
	order := handler.repository.Get(cmd.OrderId)
	order.Cancel()
	handler.repository.Update(order)
}

func (validator *CancelOrderCommandValidator) Validate(cmd CancelOrderCommand) error {
	if cmd.OrderId == "" {
		return errors.New("failed")
	}

	return nil
}
