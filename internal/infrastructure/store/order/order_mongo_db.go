package order

import (
	"context"
	"time"

	"gopkg.in/mgo.v2/bson"

	"ordercontext/internal/domain"
	"ordercontext/internal/infrastructure"
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

func FromOrder(o *domain.Order) *orderBson {
	return &orderBson{
		ID:         o.ID(),
		Status:     domain.FromStatus(o.Status()),
		CustomerID: o.CustomerID(),
		ProductID:  o.ProductID(),
		Version:    o.Version(),
	}
}

func FromBson(o *orderBson) *domain.Order {
	ord, _ := domain.NewOrder(domain.OrderID(o.ID),
		domain.CustomerID(o.CustomerID),
		domain.ProductID(o.ProductID),
		func() time.Time { return time.Now() },
		domain.ToStatus(o.Status),
		aggregate.Version(o.Version))

	ord.Clear()
	return ord
}

type MongoRepository struct {
	mStore *infrastructure.MongoStore
}

func NewMongoRepository(mongoStore *infrastructure.MongoStore) *MongoRepository {
	return &MongoRepository{mStore: mongoStore}
}

func (r *MongoRepository) GetAll(ctx context.Context) ([]*domain.Order, error) {
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

func (r *MongoRepository) Get(ctx context.Context, id string) (*domain.Order, error) {
	var (
		bsonResult = &orderBson{}
		query      = bson.M{"id": id}
	)

	if err := r.mStore.FindOne(ctx, collectionName, query, nil, bsonResult); err != nil {
		return nil, err
	}

	return FromBson(bsonResult), nil
}

func (r *MongoRepository) Update(ctx context.Context, o *domain.Order) error {
	query := bson.M{"id": o.ID(), "version": o.Version()}
	update := bson.M{"$set": bson.M{"status": o.Status(), "version": aggregate.NewVersion().String()}}

	return r.mStore.Update(ctx, collectionName, query, update)
}

func (r *MongoRepository) Create(ctx context.Context, o *domain.Order) error {
	bson := FromOrder(o)
	if bson.Version == "" {
		bson.Version = aggregate.NewVersion().String()
	}
	return r.mStore.Store(ctx, collectionName, bson)
}
