package infrastructure

import (
	"context"
	"log"
	"time"

	"orderContext/domain/customer"
	"orderContext/domain/order"
	"orderContext/domain/product"
	"orderContext/infrastructure/store"

	"gopkg.in/mgo.v2/bson"
)

type orderBson struct {
	Id          string    `bson:"id"`
	CustomerId  string    `bson:"customerId"`
	ProductId   string    `bson:"productId"`
	CreatedTime time.Time `bson:"createdTime"`
	Status      int       `bson:"status"`
}

func FromOrder(o *order.Order) *orderBson {
	return &orderBson{
		Id:         o.Id(),
		Status:     order.FromStatus(o.Status()),
		ProductId:  o.ProductId(),
		CustomerId: o.CustomerId(),
	}
}

func FromBson(o *orderBson) *order.Order {
	ord, _ := order.NewOrder(order.OrderId(o.Id),
		customer.CustomerId(o.CustomerId),
		product.ProductId(o.ProductId),
		func() time.Time { return time.Now() }, order.ToStatus(o.Status))

	ord.Clear()
	return ord
}

const collectionName = "orders"

type OrderMongoRepository struct {
	mStore         *store.MongoStore
	collectionName string
}

func NewOrderMongoRepository(mongoStore *store.MongoStore) *OrderMongoRepository {
	return &OrderMongoRepository{mStore: mongoStore, collectionName: collectionName}
}

func (r *OrderMongoRepository) GetOrders(ctx context.Context) []*order.Order {

	var result []*orderBson

	if err := r.mStore.FindAll(ctx, r.collectionName, bson.M{}, &result); err != nil {
		log.Println(err)
	}

	var orders []*order.Order

	for _, o := range result {
		orders = append(orders, FromBson(o))
	}

	return orders
}

func (r *OrderMongoRepository) Get(ctx context.Context, id string) *order.Order {
	var (
		bsonResult = &orderBson{}
		query      = bson.M{"id": id}
	)

	if err := r.mStore.FindOne(ctx, r.collectionName, query, nil, bsonResult); err != nil {
		log.Println(err)
	}

	return FromBson(bsonResult)
}

func (r *OrderMongoRepository) Update(ctx context.Context, o *order.Order) {
	var (
		query  = bson.M{"id": o.Id()}
		update = bson.M{"$set": bson.M{"status": o.Status()}}
	)

	if err := r.mStore.Update(ctx, r.collectionName, query, update); err != nil {
		log.Println(err)
	}

}

func (r *OrderMongoRepository) Create(ctx context.Context, o *order.Order) {
	bson := FromOrder(o)

	if err := r.mStore.Store(ctx, r.collectionName, bson); err != nil {
		log.Println(err)
	}
}
