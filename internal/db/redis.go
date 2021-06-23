package db

import (
	"context"
	"fmt"
	"time"

	"github.com/anonychun/go-blog-api/internal/config"
	cache "github.com/go-redis/cache/v8"
	redis "github.com/go-redis/redis/v8"
)

type RedisClient interface {
	Conn() *redis.Client
	Cache() *cache.Cache
	Close() error
}

func NewRedisClient() (RedisClient, error) {
	db := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d",
			config.Cfg().RedisHost, config.Cfg().RedisPort,
		),
		Password: config.Cfg().RedisPassword,
		DB:       config.Cfg().RedisDatabase,
		PoolSize: config.Cfg().RedisPoolSize,
	})

	err := db.Ping(context.Background()).Err()
	if err != nil {
		return nil, err
	}

	cacher := cache.New(&cache.Options{
		Redis:      db,
		LocalCache: cache.NewTinyLFU(1000, time.Minute),
	})

	return &redisClient{db, cacher}, nil
}

type redisClient struct {
	db     *redis.Client
	cacher *cache.Cache
}

func (c *redisClient) Conn() *redis.Client { return c.db }
func (c *redisClient) Cache() *cache.Cache { return c.cacher }
func (c *redisClient) Close() error        { return c.db.Close() }
