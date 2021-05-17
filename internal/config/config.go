package config

import (
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	AppPort int

	HttpRateLimitRequest int
	HttpRateLimitTime    time.Duration

	JwtSecretKey string
	JwtTTL       time.Duration

	PaginationLimit int

	MysqlUser            string
	MysqlPassword        string
	MysqlHost            string
	MysqlPort            int
	MysqlDatabase        string
	MysqlMaxIdleConns    int
	MysqlMaxOpenConns    int
	MysqlConnMaxLifetime time.Duration

	RedisPassword string
	RedisHost     string
	RedisPort     int
	RedisDatabase int
	RedisPoolSize int
	RedisTTL      time.Duration
}

func load() Config {
	fang := viper.New()
	fang.SetConfigFile(".env")
	fang.AddConfigPath(".")
	fang.AutomaticEnv()
	fang.ReadInConfig()

	return Config{
		AppPort:              fang.GetInt("APP_PORT"),
		HttpRateLimitRequest: fang.GetInt("HTTP_RATE_LIMIT_REQUEST"),
		HttpRateLimitTime:    fang.GetDuration("HTTP_RATE_LIMIT_TIME"),
		JwtSecretKey:         fang.GetString("JWT_SECRET_KEY"),
		JwtTTL:               fang.GetDuration("JWT_TTL"),
		PaginationLimit:      fang.GetInt("PAGINATION_LIMIT"),
		MysqlUser:            fang.GetString("MYSQL_USER"),
		MysqlPassword:        fang.GetString("MYSQL_PASSWORD"),
		MysqlHost:            fang.GetString("MYSQL_HOST"),
		MysqlPort:            fang.GetInt("MYSQL_PORT"),
		MysqlDatabase:        fang.GetString("MYSQL_DATABASE"),
		MysqlMaxIdleConns:    fang.GetInt("MYSQL_MAX_IDLE_CONNS"),
		MysqlMaxOpenConns:    fang.GetInt("MYSQL_MAX_OPEN_CONNS"),
		MysqlConnMaxLifetime: fang.GetDuration("MYSQL_CONN_MAX_LIFETIME"),
		RedisPassword:        fang.GetString("REDIS_PASSWORD"),
		RedisHost:            fang.GetString("REDIS_HOST"),
		RedisPort:            fang.GetInt("REDIS_PORT"),
		RedisDatabase:        fang.GetInt("REDIS_DATABASE"),
		RedisPoolSize:        fang.GetInt("REDIS_POOL_SIZE"),
		RedisTTL:             fang.GetDuration("REDIS_TTL"),
	}
}

var config = load()

func Cfg() *Config { return &config }
