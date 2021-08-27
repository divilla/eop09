package cmongo

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/net/context"
)

type Repository interface {
	List(ctx context.Context, pageNumber, pageSize int64, results interface{}) error
	FindOne(ctx context.Context, id interface{}, v interface{}) error
	CreateOne(ctx context.Context, document interface{}) error
	CreateMany(ctx context.Context, documents []interface{}) error
	UpsertOne(ctx context.Context, id interface{}, document interface{}) error
	UpdateOne(ctx context.Context, id interface{}, document interface{}) error
	DeleteOne(ctx context.Context, id interface{}) error
	DropCollection(ctx context.Context) error
}

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
		SetSort(bson.D{{"_id", 1}}).
		SetSkip(skip).
		SetLimit(pageSize))
	if err != nil {
		return err
	}

	return cur.All(ctx, results)
}

func (r *repository) FindOne(ctx context.Context, id interface{}, v interface{}) error {
	dbc, cancel := context.WithCancel(ctx)
	defer cancel()

	res := r.collection.FindOne(dbc, bson.M{"_id": id})
	if res.Err() != nil {
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

func (r *repository) UpsertOne(ctx context.Context, id interface{}, document interface{}) error {
	dbc, cancel := context.WithCancel(ctx)
	defer cancel()

	_, err := r.collection.ReplaceOne(dbc, bson.M{"_id": id}, document, options.Replace().SetUpsert(true))
	return err
}

func (r *repository) UpdateOne(ctx context.Context, id interface{}, document interface{}) error {
	dbc, cancel := context.WithCancel(ctx)
	defer cancel()

	_, err := r.collection.UpdateByID(dbc, id, document)
	return err
}

func (r *repository) DeleteOne(ctx context.Context, id interface{}) error {
	dbc, cancel := context.WithCancel(ctx)
	defer cancel()

	_, err := r.collection.DeleteOne(dbc, bson.M{"_id": id})
	return err
}

func (r *repository) DropCollection(ctx context.Context) error {
	return r.DropCollection(ctx)
}
