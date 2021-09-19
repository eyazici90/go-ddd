package application

import (
	"context"
	"time"

	"ordercontext/internal/application/behavior"
	"ordercontext/internal/application/command"
	"ordercontext/internal/application/event"
	"ordercontext/internal/domain/order"

	"github.com/eyazici90/go-mediator/pkg/mediator"
)

type OrderStore interface {
	GetAll(context.Context) ([]*order.Order, error)
	Get(ctx context.Context, id string) (*order.Order, error)
	Create(ctx context.Context, o *order.Order) error
	Update(ctx context.Context, o *order.Order) error
}

func NewMediator(store OrderStore,
	ep event.Publisher,
	timeout time.Duration) (*mediator.Mediator, error) {
	sender, err := mediator.New(
		// Behaviors
		mediator.WithBehaviourFunc(behavior.Measure),
		mediator.WithBehaviourFunc(behavior.Validate),
		mediator.WithBehaviour(behavior.NewCancellator(timeout)),
		// Handlers
		mediator.WithHandler(command.CreateOrder{}, command.NewCreateOrderHandler(store.Create)),
		mediator.WithHandler(command.PayOrder{}, command.NewPayOrderHandler(store.Get, store.Update)),
		mediator.WithHandler(command.ShipOrder{}, command.NewShipOrderHandler(store.Get, store.Update, ep)),
	)
	if err != nil {
		return nil, err
	}
	return sender, nil
}
