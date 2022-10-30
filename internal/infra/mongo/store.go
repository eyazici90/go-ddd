package mongo

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Store struct {
	db *mongo.Database
}

func NewStore(uri, dbName string, timeout time.Duration) (*Store, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	clientOpts := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
		return nil, err
	}
	return &Store{client.Database(dbName)}, nil
}

func (store *Store) Store(ctx context.Context, collection string, data interface{}) error {
	if _, err := store.db.Collection(collection).InsertOne(ctx, data); err != nil {
		return err
	}

	return nil
}

func (store *Store) Update(ctx context.Context, collection string, query, update interface{}) error {
	if _, err := store.db.Collection(collection).UpdateOne(ctx, query, update); err != nil {
		return err
	}

	return nil
}

func (store *Store) FindAll(ctx context.Context, collection string, query, result interface{}) error {
	cur, err := store.db.Collection(collection).Find(ctx, query)
	if err != nil {
		return fmt.Errorf("finding collection: %w", err)
	}
	defer func() {
		_ = cur.Close(ctx)
	}()

	if err := cur.All(ctx, result); err != nil {
		return fmt.Errorf("fetching result: %w", err)
	}

	return cur.Err()
}

func (store *Store) FindOne(
	ctx context.Context,
	collection string,
	query interface{},
	projection interface{},
	result interface{},
) error {
	return store.db.Collection(collection).
		FindOne(
			ctx,
			query,
			options.FindOne().SetProjection(projection)).Decode(result)
}
