package store

import (
	"context"
	"ordercontext/internal/domain"
	"time"

	"gopkg.in/mgo.v2/bson"

	"ordercontext/internal/infra"
	"ordercontext/pkg/aggregate"
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

func FromBson(o *orderBson) *domain.Order {
	ord, _ := domain.NewOrder(domain.OrderID(o.ID),
		domain.CustomerID(o.CustomerID),
		domain.ProductID(o.ProductID),
		time.Now,
		domain.OrderStatus(o.Status),
		aggregate.Version(o.Version))

	ord.Clear()
	return ord
}

type OrderMongoRepository struct {
	mStore *infra.MongoStore
}

func NewOrderMongoRepository(mongoStore *infra.MongoStore) *OrderMongoRepository {
	return &OrderMongoRepository{mStore: mongoStore}
}

func (r *OrderMongoRepository) GetAll(ctx context.Context) ([]*domain.Order, error) {
	var result []*orderBson

	if err := r.mStore.FindAll(ctx, collectionName, bson.M{}, &result); err != nil {
		return nil, err
	}

	var orders []*domain.Order

	for _, o := range result {
		orders = append(orders, FromBson(o))
	}

	return orders, nil
}

func (r *OrderMongoRepository) Get(ctx context.Context, id string) (*domain.Order, error) {
	var (
		bsonResult = &orderBson{}
		query      = bson.M{"id": id}
	)

	if err := r.mStore.FindOne(ctx, collectionName, query, nil, bsonResult); err != nil {
		return nil, err
	}

	return FromBson(bsonResult), nil
}

func (r *OrderMongoRepository) Update(ctx context.Context, o *domain.Order) error {
	query := bson.M{"id": o.ID(), "version": o.Version()}
	update := bson.M{"$set": bson.M{"status": o.Status(), "version": aggregate.NewVersion().String()}}

	return r.mStore.Update(ctx, collectionName, query, update)
}

func (r *OrderMongoRepository) Create(ctx context.Context, o *domain.Order) error {
	bOrder := fromOrder(o)
	if bOrder.Version == "" {
		bOrder.Version = aggregate.NewVersion().String()
	}
	return r.mStore.Store(ctx, collectionName, bOrder)
}

func fromOrder(o *domain.Order) *orderBson {
	return &orderBson{
		ID:         o.ID(),
		Status:     int(o.Status()),
		CustomerID: o.CustomerID(),
		ProductID:  o.ProductID(),
		Version:    o.Version(),
	}
}
