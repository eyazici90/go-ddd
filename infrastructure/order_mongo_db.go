package infrastructure

import (
	"context"
	"log"

	"orderContext/domain/order"
	"orderContext/infrastructure/store"

	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

type OrderMongoRepository struct {
	mStore         *store.MongoStore
	collectionName string
}

func NewOrderMongoRepository(mongoStore *store.MongoStore) *OrderMongoRepository {
	return &OrderMongoRepository{mongoStore, "orders"}
}

func (r *OrderMongoRepository) GetOrders(ctx context.Context) []*order.Order {

	var result []*order.Order

	if err := r.mStore.FindAll(ctx, r.collectionName, bson.M{}, result); err != nil {
		switch err {
		case mongo.ErrNilDocument:
			return nil
		default:
			return nil
		}
	}

	return result
}

func (r *OrderMongoRepository) Get(ctx context.Context, id string) *order.Order {
	var (
		order = &order.Order{}
		query = bson.M{"Id": id}
	)

	if err := r.mStore.FindOne(ctx, r.collectionName, query, nil, order); err != nil {
		switch err {
		case mongo.ErrNoDocuments:
			return nil
		default:
			return nil
		}
	}

	return order
}

func (r *OrderMongoRepository) Update(ctx context.Context, o *order.Order) {
	var (
		query  = bson.M{"Id": o.Id()}
		update = bson.M{"$set": bson.M{"status": o.Status()}}
	)

	if err := r.mStore.Update(ctx, r.collectionName, query, update); err != nil {
		log.Println(err)
	}

}

func (r *OrderMongoRepository) Create(ctx context.Context, o *order.Order) {

	if err := r.mStore.Store(ctx, r.collectionName, o); err != nil {
		log.Println(err)
	}
}
