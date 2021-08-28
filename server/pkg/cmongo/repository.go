package cmongo

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/net/context"
	"net/http"
)

type (
	repository struct {
		mongo      *mongo.Database
		collection *mongo.Collection
	}
)

func NewRepository(mongo *CMongo, collection string) *repository {
	return &repository{
		mongo:      mongo.Db(),
		collection: mongo.Collection(collection),
	}
}

func (r *repository) List(ctx context.Context, pageNumber, pageSize int64, results interface{}) error {
	dbc, cancel := context.WithCancel(ctx)
	defer cancel()

	skip := (pageNumber - 1) * pageSize

	cur, err := r.collection.Find(dbc, bson.D{}, options.Find().
		SetSort(bson.D{{"key", 1}}).
		SetSkip(skip).
		SetLimit(pageSize))
	if err != nil {
		return err
	}

	return cur.All(ctx, results)
}

func (r *repository) FindOne(ctx context.Context, key interface{}, v interface{}) error {
	dbc, cancel := context.WithCancel(ctx)
	defer cancel()

	res := r.collection.FindOne(dbc, bson.M{"key": key})
	err := res.Err()
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			return NewJsonError(http.StatusNotFound, "document with requested key not found")
		}
		return res.Err()
	}

	return res.Decode(v)
}

func (r *repository) CreateOne(ctx context.Context, document interface{}) error {
	dbc, cancel := context.WithCancel(ctx)
	defer cancel()

	_, err := r.collection.InsertOne(dbc, document)
	return err
}

func (r *repository) CreateMany(ctx context.Context, documents []interface{}) error {
	dbc, cancel := context.WithCancel(ctx)
	defer cancel()

	_, err := r.collection.InsertMany(dbc, documents)
	return err
}

func (r *repository) UpsertOne(ctx context.Context, key interface{}, document interface{}) error {
	dbc, cancel := context.WithCancel(ctx)
	defer cancel()

	_, err := r.collection.ReplaceOne(dbc, bson.M{"key": key}, document, options.Replace().SetUpsert(true))
	return err
}

func (r *repository) UpdateOne(ctx context.Context, key interface{}, document interface{}) error {
	dbc, cancel := context.WithCancel(ctx)
	defer cancel()

	_, err := r.collection.UpdateOne(dbc, bson.M{"key": key}, document)
	return err
}

func (r *repository) ReplaceOne(ctx context.Context, key interface{}, document interface{}) error {
	dbc, cancel := context.WithCancel(ctx)
	defer cancel()

	_, err := r.collection.ReplaceOne(dbc, bson.M{"key": key}, document)
	return err
}

func (r *repository) DeleteOne(ctx context.Context, key interface{}) error {
	dbc, cancel := context.WithCancel(ctx)
	defer cancel()

	_, err := r.collection.DeleteOne(dbc, bson.M{"key": key})
	return err
}

func (r *repository) CountAll(ctx context.Context) (int64, error) {
	return r.collection.CountDocuments(ctx, bson.D{})
}

func (r *repository) DropCollection(ctx context.Context) error {
	return r.collection.Drop(ctx)
}
