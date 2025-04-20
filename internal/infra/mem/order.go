package mem

import (
	"context"
	"sync"

	"github.com/eyazici90/go-ddd/internal/order"
	"github.com/eyazici90/go-ddd/pkg/otel"
	"go.opentelemetry.io/otel/attribute"
)

type OrderRepository struct {
	data  map[string]*order.Order
	mutex sync.RWMutex
}

func NewOrderRepository() *OrderRepository {
	return &OrderRepository{
		data: make(map[string]*order.Order),
	}
}

func (i *OrderRepository) GetAll(ctx context.Context) ([]*order.Order, error) {
	ctx, span := otel.Tracer().Start(ctx, "inmem-store-getall")
	defer span.End()

	i.mutex.RLock()
	defer i.mutex.RUnlock()

	var result []*order.Order

	for _, v := range i.data {
		result = append(result, v)
	}

	return result, nil
}

func (i *OrderRepository) Get(ctx context.Context, id string) (*order.Order, error) {
	ctx, span := otel.Tracer().Start(ctx, "inmem-store-get")
	defer span.End()

	span.SetAttributes(attribute.String("order-id", id))
	i.mutex.RLock()
	defer i.mutex.RUnlock()

	return i.data[id], nil
}

func (i *OrderRepository) Update(ctx context.Context, o *order.Order) error {
	ctx, span := otel.Tracer().Start(ctx, "inmem-store-update")
	defer span.End()

	span.SetAttributes(attribute.String("order-id", o.ID()))
	i.mutex.Lock()
	defer i.mutex.Unlock()

	i.data[o.ID()] = o
	return nil
}

func (i *OrderRepository) Create(ctx context.Context, o *order.Order) error {
	ctx, span := otel.Tracer().Start(ctx, "inmem-store-create")
	defer span.End()

	span.SetAttributes(attribute.String("order-id", o.ID()))
	i.mutex.Lock()
	defer i.mutex.Unlock()

	i.data[o.ID()] = o
	return nil
}
