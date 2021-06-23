package db

import (
	"database/sql"
	"fmt"

	"github.com/anonychun/go-blog-api/internal/config"
	_ "github.com/go-sql-driver/mysql"
)

type MysqlClient interface {
	Conn() *sql.DB
	Close() error
}

func NewMysqlClient() (MysqlClient, error) {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true",
		config.Cfg().MysqlUser,
		config.Cfg().MysqlPassword,
		config.Cfg().MysqlHost,
		config.Cfg().MysqlPort,
		config.Cfg().MysqlDatabase,
	)

	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	db.SetMaxIdleConns(config.Cfg().MysqlMaxIdleConns)
	db.SetMaxOpenConns(config.Cfg().MysqlMaxOpenConns)
	db.SetConnMaxLifetime(config.Cfg().MysqlConnMaxLifetime)

	return &mysqlClient{db}, nil
}

type mysqlClient struct {
	db *sql.DB
}

func (c *mysqlClient) Conn() *sql.DB { return c.db }
func (c *mysqlClient) Close() error  { return c.db.Close() }
