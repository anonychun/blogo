package config

import (
	"os"
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

	PostgresUser            string
	PostgresPassword        string
	PostgresHost            string
	PostgresPort            int
	PostgresDatabase        string
	PostgresMaxIdleConns    int
	PostgresMaxOpenConns    int
	PostgresConnMaxLifetime time.Duration

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
	configLocation, available := os.LookupEnv("CONFIG_LOCATION")
	if available {
		fang.AddConfigPath(configLocation)
	}

	fang.AutomaticEnv()
	fang.ReadInConfig()

	return Config{
		AppPort:                 fang.GetInt("APP_PORT"),
		HttpRateLimitRequest:    fang.GetInt("HTTP_RATE_LIMIT_REQUEST"),
		HttpRateLimitTime:       fang.GetDuration("HTTP_RATE_LIMIT_TIME"),
		JwtSecretKey:            fang.GetString("JWT_SECRET_KEY"),
		JwtTTL:                  fang.GetDuration("JWT_TTL"),
		PaginationLimit:         fang.GetInt("PAGINATION_LIMIT"),
		PostgresUser:            fang.GetString("POSTGRES_USER"),
		PostgresPassword:        fang.GetString("POSTGRES_PASSWORD"),
		PostgresHost:            fang.GetString("POSTGRES_HOST"),
		PostgresPort:            fang.GetInt("POSTGRES_PORT"),
		PostgresDatabase:        fang.GetString("POSTGRES_DATABASE"),
		PostgresMaxIdleConns:    fang.GetInt("POSTGRES_MAX_IDLE_CONNS"),
		PostgresMaxOpenConns:    fang.GetInt("POSTGRES_MAX_OPEN_CONNS"),
		PostgresConnMaxLifetime: fang.GetDuration("POSTGRES_CONN_MAX_LIFETIME"),
		RedisPassword:           fang.GetString("REDIS_PASSWORD"),
		RedisHost:               fang.GetString("REDIS_HOST"),
		RedisPort:               fang.GetInt("REDIS_PORT"),
		RedisDatabase:           fang.GetInt("REDIS_DATABASE"),
		RedisPoolSize:           fang.GetInt("REDIS_POOL_SIZE"),
		RedisTTL:                fang.GetDuration("REDIS_TTL"),
	}
}

var config = load()

func Cfg() *Config { return &config }
