package app

import (
	"fmt"
	"time"

	"github.com/eyazici90/go-ddd/internal/app/create"
	"github.com/eyazici90/go-ddd/internal/app/feature"
	"github.com/eyazici90/go-ddd/internal/app/pay"
	"github.com/eyazici90/go-ddd/internal/app/ship"
	"github.com/eyazici90/go-ddd/internal/infra/behavior"
	"github.com/eyazici90/go-mediator/mediator"
)

type OrderStore interface {
	create.OrderCreator
	feature.OrderGetter
	feature.OrderUpdater
}

func NewMediator(store OrderStore,
	ep feature.EventPublisher,
	timeout time.Duration,
) (*mediator.Mediator, error) {
	m, err := mediator.New(
		// Behaviors
		mediator.WithBehaviourFunc(behavior.Measure),
		mediator.WithBehaviourFunc(behavior.Validate),
		mediator.WithBehaviour(behavior.NewCancellator(timeout)),
		// Handlers
		mediator.WithHandler(create.OrderCommand{}, create.NewOrderCommandHandler(store)),
		mediator.WithHandler(pay.OrderCommand{}, pay.NewOrderCommandHandler(store, store)),
		mediator.WithHandler(ship.OrderCommand{}, ship.NewOrderCommandHandler(store, store, ep)),
	)
	if err != nil {
		return nil, fmt.Errorf("create mediator: %w", err)
	}
	return m, nil
}
