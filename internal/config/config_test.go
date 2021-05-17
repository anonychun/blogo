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
	assert.NotEmpty(t, Cfg().MysqlUser, "MYSQL_USER")
	assert.NotEmpty(t, Cfg().MysqlPassword, "MYSQL_PASSWORD")
	assert.NotEmpty(t, Cfg().MysqlHost, "MYSQL_HOST")
	assert.NotZero(t, Cfg().MysqlPort, "MYSQL_PORT")
	assert.NotEmpty(t, Cfg().MysqlDatabase, "MYSQL_DATABASE")
	assert.NotZero(t, Cfg().MysqlMaxIdleConns, "MYSQL_MAX_IDLE_CONNS")
	assert.NotZero(t, Cfg().MysqlMaxOpenConns, "MYSQL_MAX_OPEN_CONNS")
	assert.NotEmpty(t, Cfg().MysqlConnMaxLifetime, "MYSQL_CONN_MAX_LIFETIME")
	assert.NotEmpty(t, Cfg().RedisPassword, "REDIS_PASSWORD")
	assert.NotEmpty(t, Cfg().RedisHost, "REDIS_HOST")
	assert.NotZero(t, Cfg().RedisPort, "REDIS_PORT")
	assert.GreaterOrEqual(t, Cfg().RedisDatabase, 0, "REDIS_DATABASE")
	assert.NotZero(t, Cfg().RedisPoolSize, "REDIS_POOL_SIZE")
	assert.NotEmpty(t, Cfg().RedisTTL, "REDIS_TTL")
}
