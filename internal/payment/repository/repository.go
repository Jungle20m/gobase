package repository

import (
	"context"
	"gorm.io/gorm"
)

const TransactionContextKey = "TransactionContextKey"

type Repo struct {
	db *gorm.DB
}

type Option func(repo *Repo)

func WithGorm(db *gorm.DB) Option {
	return func(repo *Repo) {
		repo.db = db
	}
}

func NewInstance(opts ...Option) *Repo {
	repo := &Repo{}

	for _, opt := range opts {
		opt(repo)
	}

	return repo
}

func (r *Repo) getDB(ctx context.Context) *gorm.DB {
	conn, ok := ctx.Value(TransactionContextKey).(*gorm.DB)
	if ok {
		return conn
	}
	return r.db
}

func (r *Repo) WithTx(ctx context.Context, fn func(txContext context.Context) error) error {
	tx := r.db.Begin()
	defer tx.Rollback()

	c := context.WithValue(ctx, TransactionContextKey, tx)
	err := fn(c)

	if err != nil {
		return err
	}
	tx.Commit()
	return nil
}
