package config

import (
	"strings"
	"sync"
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

var once sync.Once
var config Config

func load() Config {
	fang := viper.New()
	fang.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	fang.AutomaticEnv()

	return Config{
		AppPort:              fang.GetInt("app.port"),
		HttpRateLimitRequest: fang.GetInt("http.rate.limit.request"),
		HttpRateLimitTime:    fang.GetDuration("http.rate.limit.time"),
		JwtSecretKey:         fang.GetString("jwt.secret.key"),
		JwtTTL:               fang.GetDuration("jwt.ttl"),
		PaginationLimit:      fang.GetInt("pagination.limit"),
		MysqlUser:            fang.GetString("mysql.user"),
		MysqlPassword:        fang.GetString("mysql.password"),
		MysqlHost:            fang.GetString("mysql.host"),
		MysqlPort:            fang.GetInt("mysql.port"),
		MysqlDatabase:        fang.GetString("mysql.database"),
		MysqlMaxIdleConns:    fang.GetInt("mysql.max.idle.conns"),
		MysqlMaxOpenConns:    fang.GetInt("mysql.max.open.conns"),
		MysqlConnMaxLifetime: fang.GetDuration("mysql.conn.max.lifetime"),
		RedisPassword:        fang.GetString("redis.password"),
		RedisHost:            fang.GetString("redis.host"),
		RedisPort:            fang.GetInt("redis.port"),
		RedisDatabase:        fang.GetInt("redis.database"),
		RedisPoolSize:        fang.GetInt("redis.pool.size"),
		RedisTTL:             fang.GetDuration("redis.ttl"),
	}
}

func Cfg() *Config {
	once.Do(func() {
		config = load()
	})
	return &config
}
