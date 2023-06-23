package gorm

import (
	"database/sql"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

type ConnectionType string

const (
	Mysql    string = "mysql"
	Postgres string = "postgres"
)

const (
	_defaultMaxOpenConnection     = 100
	_defaultMaxIdleConnection     = 10
	_defaultConnectionMaxLifetime = time.Hour
)

type Database struct {
	maxOpenConnection     int
	maxIdleConnection     int
	connectionMaxLifetime time.Duration
	connectionType        string

	conn *gorm.DB
}

func NewInstance(dsn string, opts ...Option) (*Database, error) {
	db := &Database{
		maxOpenConnection:     _defaultMaxOpenConnection,
		maxIdleConnection:     _defaultMaxIdleConnection,
		connectionMaxLifetime: _defaultConnectionMaxLifetime,
		connectionType:        Mysql,
	}

	for _, opt := range opts {
		opt(db)
	}

	switch db.connectionType {
	case Mysql:
		conn, err := sql.Open("mysql", dsn)
		if err != nil {
			return nil, err
		}
		conn.SetMaxOpenConns(db.maxOpenConnection)
		conn.SetMaxIdleConns(db.maxIdleConnection)
		conn.SetConnMaxLifetime(db.connectionMaxLifetime)

		gormDB, err := gorm.Open(mysql.New(mysql.Config{
			Conn: conn,
		}), &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)})
		if err != nil {
			return nil, err
		}

		db.conn = gormDB
		return db, nil
	case Postgres:
		gormDB, err := gorm.Open(postgres.New(postgres.Config{
			DSN: dsn,
		}), &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)})
		if err != nil {
			return nil, err
		}
		db.conn = gormDB
		return db, nil
	default:
		return nil, fmt.Errorf("mgorm only support mysql or postgres")
	}
}

func (db *Database) GetDB() *gorm.DB {
	return db.conn
}
