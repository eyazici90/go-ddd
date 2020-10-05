package application

import (
	"orderContext/application/behaviour"
	"orderContext/application/command"
	"orderContext/domain/order"
	"orderContext/infrastructure"

	"github.com/eyazici90/go-mediator"
)

func NewMediator(r order.Repository, e infrastructure.EventPublisher, timeout int) mediator.Mediator {
	m, _ := mediator.New().
		UseBehaviour(behaviour.NewMeasurer()).
		UseBehaviour(behaviour.NewLogger()).
		UseBehaviour(behaviour.NewValidator()).
		UseBehaviour(behaviour.NewCancellator(timeout)).
		UseBehaviour(behaviour.NewRetrier()).
		RegisterHandler(command.CreateOrderCommand{}, command.NewCreateOrderCommandHandler(r.Create)).
		RegisterHandler(command.PayOrderCommand{}, command.NewPayOrderCommandHandler(r.Get, r.Update)).
		RegisterHandler(command.ShipOrderCommand{}, command.NewShipOrderCommandHandler(r, e)).
		Build()

	return m

}
