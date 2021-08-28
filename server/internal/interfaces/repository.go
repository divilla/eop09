package interfaces

import "golang.org/x/net/context"

type Repository interface {
	List(ctx context.Context, pageNumber, pageSize int64, results interface{}) error
	FindOne(ctx context.Context, id interface{}, v interface{}) error
	CreateOne(ctx context.Context, document interface{}) error
	CreateMany(ctx context.Context, documents []interface{}) error
	UpsertOne(ctx context.Context, id interface{}, document interface{}) error
	UpdateOne(ctx context.Context, id interface{}, document interface{}) error
	ReplaceOne(ctx context.Context, id interface{}, document interface{}) error
	DeleteOne(ctx context.Context, id interface{}) error
	CountAll(ctx context.Context) (int64, error)
	DropCollection(ctx context.Context) error
}

