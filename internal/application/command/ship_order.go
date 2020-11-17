package command

import (
	"context"

	"ordercontext/internal/domain/order"
	"ordercontext/internal/infrastructure"

	"github.com/eyazici90/go-mediator"
)

type ShipOrderCommand struct {
	OrderId string `validate:"required,min=10"`
}

func (ShipOrderCommand) Key() string { return "ShipOrderCommand" }

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

func (handler ShipOrderCommandHandler) Handle(ctx context.Context, request mediator.Message) error {
	cmd := request.(ShipOrderCommand)
	order, err := handler.repository.Get(ctx, cmd.OrderId)
	if err != nil {
		return nil
	}

	err = order.Ship()

	if err != nil {
		return err
	}

	handler.repository.Update(ctx, order)

	handler.eventPublisher.PublishAll(order.Events())

	return nil
}
