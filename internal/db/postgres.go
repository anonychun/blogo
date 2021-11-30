package db

import (
	"context"
	"fmt"

	"github.com/anonychun/go-blog-api/internal/config"
	pgx "github.com/jackc/pgx/v4"
)

type PostgresClient interface {
	Conn() *pgx.Conn
	Close() error
}

func NewPostgresClientContext(ctx context.Context) (PostgresClient, error) {
	dataSourceName := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		config.Cfg().PostgresUser,
		config.Cfg().PostgresPassword,
		config.Cfg().PostgresHost,
		config.Cfg().PostgresPort,
		config.Cfg().PostgresDatabase,
	)

	db, err := pgx.Connect(ctx, dataSourceName)
	if err != nil {
		panic(err)
	}

	err = db.Ping(ctx)
	if err != nil {
		return nil, err
	}

	return &postgresClient{db}, nil
}

func NewPostgresClient() (PostgresClient, error) {
	return NewPostgresClientContext(context.Background())
}

type postgresClient struct {
	db *pgx.Conn
}

func (c *postgresClient) Conn() *pgx.Conn { return c.db }
func (c *postgresClient) Close() error    { return c.db.Close(context.Background()) }
