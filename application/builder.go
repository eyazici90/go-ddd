package application

import (
	"ordercontext/application/behaviour"
	"ordercontext/application/command"
	"ordercontext/domain/order"
	"ordercontext/infrastructure"

	"github.com/eyazici90/go-mediator"
)

func NewMediator(r order.Repository, e infrastructure.EventPublisher, timeout int) mediator.Mediator {
	m, _ := mediator.New().
		Use(behaviour.Measure).
		Use(behaviour.Log).
		Use(behaviour.Validate).
		UseBehaviour(behaviour.NewCancellator(timeout)).
		Use(behaviour.Retry).
		RegisterHandler(command.CreateOrderCommand{}, command.NewCreateOrderCommandHandler(r.Create)).
		RegisterHandler(command.PayOrderCommand{}, command.NewPayOrderCommandHandler(r.Get, r.Update)).
		RegisterHandler(command.ShipOrderCommand{}, command.NewShipOrderCommandHandler(r, e)).
		Build()

	return m

}
