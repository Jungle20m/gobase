package gorm

import "time"

type Option func(database *Database)

func WithMaxOpenConnection(size int) Option {
	return func(db *Database) {
		db.maxOpenConnection = size
	}
}

func WithMaxIdleConnection(size int) Option {
	return func(db *Database) {
		db.maxIdleConnection = size
	}
}

func WithConnectionMaxLifetime(duration time.Duration) Option {
	return func(db *Database) {
		db.connectionMaxLifetime = duration
	}
}

func WithConnectionType(connectionType string) Option {
	return func(db *Database) {
		db.connectionType = connectionType
	}
}
