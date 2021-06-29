package application

import (
	"time"

	"ordercontext/internal/application/behavior"
	"ordercontext/internal/application/command"
	"ordercontext/internal/application/event"
	"ordercontext/internal/domain"
	"ordercontext/pkg/must"

	"github.com/eyazici90/go-mediator/mediator"
)

func NewMediator(repository domain.OrderRepository,
	ep event.Publisher,
	timeout time.Duration) mediator.Sender {
	sender, err := mediator.NewContext().
		Use(behavior.Measure).
		Use(behavior.Log).
		Use(behavior.Validate).
		UseBehaviour(behavior.NewCancellator(timeout)).
		Use(behavior.Retry).
		RegisterHandler(command.CreateOrderCommand{}, command.NewCreateOrderCommandHandler(repository.Create)).
		RegisterHandler(command.PayOrderCommand{}, command.NewPayOrderCommandHandler(repository.Get, repository.Update)).
		RegisterHandler(command.ShipOrderCommand{}, command.NewShipOrderCommandHandler(repository, ep)).
		Build()

	must.NotFail(err)
	return sender
}
