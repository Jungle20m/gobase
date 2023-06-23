package utils

import (
	"gobase/logger"
	"gorm.io/gorm"
)

type IUtils interface {
	GetDB() *gorm.DB
	GetLogger() logger.ILogger
}

type utils struct {
	db *gorm.DB
	l  logger.ILogger
}

func NewInstance(db *gorm.DB, logger logger.ILogger) *utils {
	return &utils{
		db: db,
		l:  logger,
	}
}

func (u *utils) GetDB() *gorm.DB {
	return u.db
}

func (u *utils) GetLogger() logger.ILogger {
	return u.l
}
