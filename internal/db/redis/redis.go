package redis

import (
	"context"
	"fmt"
	"time"

	"github.com/anonychun/go-blog-api/internal/config"
	cache "github.com/go-redis/cache/v8"
	redis "github.com/go-redis/redis/v8"
)

type Client interface {
	Conn() *redis.Client
	Cache() *cache.Cache
	Close() error
}

func NewClientContext(ctx context.Context) (Client, error) {
	db := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d",
			config.Cfg().RedisHost, config.Cfg().RedisPort,
		),
		Password: config.Cfg().RedisPassword,
		DB:       config.Cfg().RedisDatabase,
		PoolSize: config.Cfg().RedisPoolSize,
	})

	err := db.Ping(ctx).Err()
	if err != nil {
		return nil, err
	}

	cacher := cache.New(&cache.Options{
		Redis:      db,
		LocalCache: cache.NewTinyLFU(1000, time.Minute),
	})

	return &client{db, cacher}, nil
}

func NewClient() (Client, error) {
	return NewClientContext(context.Background())
}

type client struct {
	db     *redis.Client
	cacher *cache.Cache
}

func (c *client) Conn() *redis.Client { return c.db }
func (c *client) Cache() *cache.Cache { return c.cacher }
func (c *client) Close() error        { return c.db.Close() }
