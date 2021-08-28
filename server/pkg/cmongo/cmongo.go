package cmongo

import (
	"context"
	interfaces2 "github.com/divilla/eop09/server/internal/interfaces"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

type (
	CMongo struct {
		client *mongo.Client
		db     *mongo.Database
	}
)

func Init(dsn string, logger interfaces2.Logger) *CMongo {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().
		ApplyURI(dsn).
		SetMinPoolSize(10).
		SetMaxPoolSize(100))
	if err != nil {
		panic(err)
	}

	//defer func() {
	//	if err = client.Disconnect(ctx); err != nil {
	//		panic(err)
	//	}
	//}()

	err = client.Ping(context.Background(), readpref.Primary())
	if err != nil {
		logger.Fatalf("unable to start mongo: %s", err)
	}

	db := client.Database("eop09")

	createCollectionAndIndex(ctx, db, "port", "key", 1)

	return &CMongo{
		client: client,
		db:     db,
	}
}

func (c *CMongo) Client() *mongo.Client {
	return c.client
}

func (c *CMongo) Db() *mongo.Database {
	return c.db
}

func (c *CMongo) Collection(name string, opts ...*options.CollectionOptions) *mongo.Collection {
	return c.db.Collection(name, opts...)
}

func createCollectionAndIndex(ctx context.Context, db *mongo.Database, collection string, field string, direction int) {
	var portExists bool
	collections, err := db.ListCollectionNames(ctx, bson.D{})
	if err != nil {
		panic(err)
	}
	for _, v := range collections {
		if v == collection {
			portExists = true
		}
	}

	if !portExists {
		err = db.CreateCollection(ctx, collection)
		if err != nil {
			panic(err)
		}
	}

	_, err = db.Collection(collection).Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys:    bson.M{field: direction},
		Options: options.Index().SetUnique(true),
	})
	if err != nil {
		panic(err)
	}
}
