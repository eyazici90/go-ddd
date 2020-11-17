package infrastructure

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoStore struct {
	db *mongo.Database
}

func NewMongoStore(uri, dbName string, timeout time.Duration) *MongoStore {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	clientOpts := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(ctx, clientOpts)

	if err != nil {
		panic(err)
	}
	return &MongoStore{client.Database(dbName)}
}

func (store *MongoStore) Store(ctx context.Context, collection string, data interface{}) error {
	if _, err := store.db.Collection(collection).InsertOne(ctx, data); err != nil {
		return err
	}

	return nil
}

func (store *MongoStore) Update(ctx context.Context, collection string, query interface{}, update interface{}) error {
	if _, err := store.db.Collection(collection).UpdateOne(ctx, query, update); err != nil {
		return err
	}

	return nil
}

func (store *MongoStore) FindAll(ctx context.Context, collection string, query interface{}, result interface{}) error {
	cur, err := store.db.Collection(collection).Find(ctx, query)
	if err != nil {
		return err
	}

	defer cur.Close(ctx)
	if err = cur.All(ctx, result); err != nil {
		return err
	}

	if err := cur.Err(); err != nil {
		return err
	}

	return nil
}

func (store *MongoStore) FindOne(
	ctx context.Context,
	collection string,
	query interface{},
	projection interface{},
	result interface{}) error {
	if err := store.db.Collection(collection).
		FindOne(
			ctx,
			query,
			options.FindOne().SetProjection(projection)).Decode(result); err != nil {
		return err
	}

	return nil
}
