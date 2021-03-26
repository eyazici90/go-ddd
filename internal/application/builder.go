package application

import (
	"ordercontext/internal/application/behaviour"
	"ordercontext/internal/application/command"
	"ordercontext/internal/domain"
	"ordercontext/internal/infrastructure"

	"github.com/eyazici90/go-mediator/mediator"
)

func NewMediator(repository domain.OrderRepository,
	ePublisher infrastructure.EventPublisher,
	timeout int) mediator.Sender {
	sender, _ := mediator.NewContext().
		Use(behaviour.Measure).
		Use(behaviour.Log).
		Use(behaviour.Validate).
		UseBehaviour(behaviour.NewCancellator(timeout)).
		Use(behaviour.Retry).
		RegisterHandler(command.CreateOrderCommand{}, command.NewCreateOrderCommandHandler(repository.Create)).
		RegisterHandler(command.PayOrderCommand{}, command.NewPayOrderCommandHandler(repository.Get, repository.Update)).
		RegisterHandler(command.ShipOrderCommand{}, command.NewShipOrderCommandHandler(repository, ePublisher)).
		Build()

	return sender

}
