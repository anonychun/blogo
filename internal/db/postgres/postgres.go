package postgres

import (
	"context"
	"fmt"

	"github.com/anonychun/go-blog-api/internal/config"
	pgx "github.com/jackc/pgx/v4"
)

type Client interface {
	Conn() *pgx.Conn
	Close() error
}

func NewClientContext(ctx context.Context) (Client, error) {
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

	return &client{db}, nil
}

func NewClient() (Client, error) {
	return NewClientContext(context.Background())
}

type client struct {
	db *pgx.Conn
}

func (c *client) Conn() *pgx.Conn { return c.db }
func (c *client) Close() error    { return c.db.Close(context.Background()) }
