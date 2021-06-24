package application

import (
	"ordercontext/internal/application/behavior"
	"ordercontext/internal/application/command"
	"ordercontext/internal/domain"
	"ordercontext/internal/infrastructure"
	"time"

	"github.com/eyazici90/go-mediator/mediator"
)

func NewMediator(repository domain.OrderRepository,
	ePublisher infrastructure.EventPublisher,
	timeout time.Duration) mediator.Sender {
	sender, _ := mediator.NewContext().
		Use(behavior.Measure).
		Use(behavior.Log).
		Use(behavior.Validate).
		UseBehaviour(behavior.NewCancellator(timeout)).
		Use(behavior.Retry).
		RegisterHandler(command.CreateOrderCommand{}, command.NewCreateOrderCommandHandler(repository.Create)).
		RegisterHandler(command.PayOrderCommand{}, command.NewPayOrderCommandHandler(repository.Get, repository.Update)).
		RegisterHandler(command.ShipOrderCommand{}, command.NewShipOrderCommandHandler(repository, ePublisher)).
		Build()

	return sender
}
