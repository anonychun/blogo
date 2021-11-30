package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfig(t *testing.T) {
	assert.NotZero(t, Cfg().AppPort, "APP_PORT")
	assert.NotZero(t, Cfg().HttpRateLimitRequest, "HTTP_RATE_LIMIT_REQUEST")
	assert.NotEmpty(t, Cfg().HttpRateLimitTime, "HTTP_RATE_LIMIT_TIME")
	assert.NotEmpty(t, Cfg().JwtSecretKey, "JWT_SECRET_KEY")
	assert.NotEmpty(t, Cfg().JwtTTL, "JWT_TTL")
	assert.NotZero(t, Cfg().PaginationLimit, "PAGINATION_LIMIT")
	assert.NotEmpty(t, Cfg().PostgresUser, "POSTGRES_USER")
	assert.NotEmpty(t, Cfg().PostgresPassword, "POSTGRES_PASSWORD")
	assert.NotEmpty(t, Cfg().PostgresHost, "POSTGRES_HOST")
	assert.NotZero(t, Cfg().PostgresPort, "POSTGRES_PORT")
	assert.NotEmpty(t, Cfg().PostgresDatabase, "POSTGRES_DATABASE")
	assert.NotZero(t, Cfg().PostgresMaxIdleConns, "POSTGRES_MAX_IDLE_CONNS")
	assert.NotZero(t, Cfg().PostgresMaxOpenConns, "POSTGRES_MAX_OPEN_CONNS")
	assert.NotEmpty(t, Cfg().PostgresConnMaxLifetime, "POSTGRES_CONN_MAX_LIFETIME")
	assert.NotEmpty(t, Cfg().RedisPassword, "REDIS_PASSWORD")
	assert.NotEmpty(t, Cfg().RedisHost, "REDIS_HOST")
	assert.NotZero(t, Cfg().RedisPort, "REDIS_PORT")
	assert.GreaterOrEqual(t, Cfg().RedisDatabase, 0, "REDIS_DATABASE")
	assert.NotZero(t, Cfg().RedisPoolSize, "REDIS_POOL_SIZE")
	assert.NotEmpty(t, Cfg().RedisTTL, "REDIS_TTL")
}
