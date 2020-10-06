package command

import (
	"context"
	"orderContext/domain/order"
	"orderContext/infrastructure"
)

type ShipOrderCommand struct {
	OrderId string `validate:"required,min=10"`
}

type ShipOrderCommandHandler struct {
	repository     order.Repository
	eventPublisher infrastructure.EventPublisher
}

func NewShipOrderCommandHandler(r order.Repository, e infrastructure.EventPublisher) ShipOrderCommandHandler {
	return ShipOrderCommandHandler{
		repository:     r,
		eventPublisher: e,
	}
}

func (handler ShipOrderCommandHandler) Handle(ctx context.Context, request interface{}) error {
	cmd := request.(ShipOrderCommand)
	order := handler.repository.Get(ctx, cmd.OrderId)

	err := order.Ship()

	if err != nil {
		return err
	}
	handler.repository.Update(ctx, order)

	handler.eventPublisher.PublishAll(order.Events())

	return nil
}
