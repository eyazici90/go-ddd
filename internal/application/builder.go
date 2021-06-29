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
	sender, err := mediator.NewContext(
		mediator.WithBehaviourFunc(behavior.Measure),
		mediator.WithBehaviourFunc(behavior.Validate),
		mediator.WithBehaviour(behavior.NewCancellator(timeout)),
		mediator.WithBehaviourFunc(behavior.Retry),
		mediator.WithHandler(command.CreateOrderCommand{}, command.NewCreateOrderCommandHandler(repository.Create)),
		mediator.WithHandler(command.PayOrderCommand{}, command.NewPayOrderCommandHandler(repository.Get, repository.Update)),
		mediator.WithHandler(command.ShipOrderCommand{}, command.NewShipOrderCommandHandler(repository, ep)),
	).Build()

	must.NotFail(err)
	return sender
}
