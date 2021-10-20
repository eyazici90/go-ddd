package application

import (
	"time"

	"ordercontext/internal/application/behavior"
	"ordercontext/internal/application/command"
	"ordercontext/internal/application/event"

	"github.com/eyazici90/go-mediator/pkg/mediator"
	"github.com/pkg/errors"
)

type OrderStore interface {
	command.OrderCreator
	command.OrderGetter
	command.OrderUpdater
}

func NewMediator(store OrderStore,
	ep event.Publisher,
	timeout time.Duration) (*mediator.Mediator, error) {
	m, err := mediator.New(
		// Behaviors
		mediator.WithBehaviourFunc(behavior.Measure),
		mediator.WithBehaviourFunc(behavior.Validate),
		mediator.WithBehaviour(behavior.NewCancellator(timeout)),
		// Handlers
		mediator.WithHandler(command.CreateOrder{}, command.NewCreateOrderHandler(store)),
		mediator.WithHandler(command.PayOrder{}, command.NewPayOrderHandler(store, store)),
		mediator.WithHandler(command.ShipOrder{}, command.NewShipOrderHandler(store, store, ep)),
	)
	if err != nil {
		return nil, errors.Wrap(err, "create mediator")
	}
	return m, nil
}
