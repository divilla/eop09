package cmongo

import (
	"context"
	"github.com/divilla/eop09/server/interfaces"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

type (
	CMongo struct {
		client         *mongo.Client
		db             *mongo.Database
	}
)

func Init(dsn string, logger interfaces.Logger) *CMongo {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	logger.Info(dsn)
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

	return &CMongo{
		client:         client,
		db:             db,
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
