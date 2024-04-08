package app

import (
	"fmt"
	"time"

	"github.com/eyazici90/go-ddd/internal/app/behavior"
	"github.com/eyazici90/go-ddd/internal/app/command"
	"github.com/eyazici90/go-ddd/internal/app/event"
	"github.com/eyazici90/go-ddd/pkg/otel"
	"github.com/eyazici90/go-mediator/mediator"
)

type OrderStore interface {
	command.OrderCreator
	command.OrderGetter
	command.OrderUpdater
}

func NewMediator(store OrderStore,
	ep event.Publisher,
	otl *otel.OTel,
	timeout time.Duration,
) (*mediator.Mediator, error) {
	m, err := mediator.New(
		// Behaviors
		mediator.WithBehaviourFunc(behavior.Measure(otl)),
		mediator.WithBehaviourFunc(behavior.Validate),
		mediator.WithBehaviour(behavior.NewCancellator(timeout)),
		// Handlers
		mediator.WithHandler(command.CreateOrder{}, command.NewCreateOrderHandler(store)),
		mediator.WithHandler(command.PayOrder{}, command.NewPayOrderHandler(store, store)),
		mediator.WithHandler(command.ShipOrder{}, command.NewShipOrderHandler(store, store, ep)),
	)
	if err != nil {
		return nil, fmt.Errorf("create mediator: %w", err)
	}
	return m, nil
}
