package pgpool

import (
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/labstack/echo/v4"
	"golang.org/x/net/context"
)

type PGPool struct {
	pool   *pgxpool.Pool
	logger echo.Logger
}

func Init(dsn string, logger echo.Logger) *PGPool {
	config, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		panic(fmt.Errorf("invalid connection string: '%s'", err))
	}

	config.MinConns = 10
	config.MaxConns = 90
	pool, err := pgxpool.ConnectConfig(context.Background(), config)
	if err != nil {
		panic(fmt.Errorf("unable to connect database pool: '%s'", err))
	}

	return &PGPool{
		pool:   pool,
		logger: logger,
	}
}

func (db *PGPool) Pool() *pgxpool.Pool {
	return db.pool
}

func (db *PGPool) Conn(ctx context.Context) (*pgxpool.Conn, error) {
	conn, err := db.pool.Acquire(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to acquire database connection from pool: '%s'", err)
	}

	return conn, nil
}

func (db *PGPool) Close() {
	db.pool.Close()
}

