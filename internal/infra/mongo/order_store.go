package mongo

import (
	"context"
	"time"

	"github.com/eyazici90/go-ddd/pkg/aggregate"
	"gopkg.in/mgo.v2/bson"
)

const collectionName = "orders"

type orderBson struct {
	ID          string    `bson:"id"`
	CustomerID  string    `bson:"customerId"`
	ProductID   string    `bson:"productId"`
	CreatedTime time.Time `bson:"createdTime"`
	Status      int       `bson:"status"`
	Version     string    `bson:"version"`
}

func FromOrderBson(o *orderBson) *order.Order {
	ord, _ := order.NewOrder(order.OrderID(o.ID),
		order.CustomerID(o.CustomerID),
		order.ProductID(o.ProductID),
		time.Now,
		order.OrderStatus(o.Status),
		aggregate.Version(o.Version))

	ord.Clear()
	return ord
}

type OrderRepository struct {
	mStore *Store
}

func NewOrderRepository(mongoStore *Store) *OrderRepository {
	return &OrderRepository{mStore: mongoStore}
}

func (r *OrderRepository) GetAll(ctx context.Context) ([]*order.Order, error) {
	var result []*orderBson
	if err := r.mStore.FindAll(ctx, collectionName, bson.M{}, &result); err != nil {
		return nil, err
	}

	var orders []*order.Order
	for _, o := range result {
		orders = append(orders, FromOrderBson(o))
	}

	return orders, nil
}

func (r *OrderRepository) Get(ctx context.Context, id string) (*order.Order, error) {
	var (
		bsonResult = &orderBson{}
		query      = bson.M{"id": id}
	)

	if err := r.mStore.FindOne(ctx, collectionName, query, nil, bsonResult); err != nil {
		return nil, err
	}

	return FromOrderBson(bsonResult), nil
}

func (r *OrderRepository) Update(ctx context.Context, o *order.Order) error {
	query := bson.M{"id": o.ID(), "version": o.Version()}
	update := bson.M{"$set": bson.M{"status": o.Status(), "version": aggregate.NewVersion().String()}}

	return r.mStore.Update(ctx, collectionName, query, update)
}

func (r *OrderRepository) Create(ctx context.Context, o *order.Order) error {
	bOrder := fromOrder(o)
	if bOrder.Version == "" {
		bOrder.Version = aggregate.NewVersion().String()
	}

	return r.mStore.Store(ctx, collectionName, bOrder)
}

func fromOrder(o *order.Order) *orderBson {
	return &orderBson{
		ID:         o.ID(),
		Status:     int(o.Status()),
		CustomerID: o.CustomerID(),
		ProductID:  o.ProductID(),
		Version:    o.Version(),
	}
}
