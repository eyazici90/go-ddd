package command

import (
	"context"
	"orderContext/domain/order"
)

type ShipOrderCommand struct {
	OrderId string `validate:"required,min=10"`
}

type ShipOrderCommandHandler struct {
	repository order.OrderRepository
}

func NewShipOrderCommandHandler(r order.OrderRepository) ShipOrderCommandHandler {
	return ShipOrderCommandHandler{repository: r}
}

func (handler ShipOrderCommandHandler) Handle(ctx context.Context, cmd ShipOrderCommand) error {
	order := handler.repository.Get(ctx, cmd.OrderId)
	err := order.Ship()
	if err != nil {
		return err
	}
	handler.repository.Update(ctx, order)
	return nil
}
