package application

import (
	"context"
	"time"

	"ordercontext/internal/application/behavior"
	"ordercontext/internal/application/command"
	"ordercontext/internal/application/event"
	"ordercontext/internal/domain"
	"ordercontext/pkg/must"

	"github.com/eyazici90/go-mediator/mediator"
)

type OrderStore interface {
	GetAll(context.Context) ([]*domain.Order, error)
	Get(ctx context.Context, id string) (*domain.Order, error)
	Create(ctx context.Context, o *domain.Order) error
	Update(ctx context.Context, o *domain.Order) error
}

func NewMediator(store OrderStore,
	ep event.Publisher,
	timeout time.Duration) mediator.Sender {
	sender, err := mediator.NewContext(
		// Behaviours
		mediator.WithBehaviourFunc(behavior.Measure),
		mediator.WithBehaviourFunc(behavior.Validate),
		mediator.WithBehaviour(behavior.NewCancellator(timeout)),
		mediator.WithBehaviourFunc(behavior.Retry),
		// Handlers
		mediator.WithHandler(command.CreateOrderCommand{}, command.NewCreateOrderCommandHandler(store.Create)),
		mediator.WithHandler(command.PayOrderCommand{}, command.NewPayOrderCommandHandler(store.Get, store.Update)),
		mediator.WithHandler(command.ShipOrderCommand{}, command.NewShipOrderCommandHandler(store.Get, store.Update, ep)),
	).Build()

	must.NotFail(err)
	return sender
}
