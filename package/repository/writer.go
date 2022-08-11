package repository

import "gorm.io/gorm"

type WriterDB struct {
	db *gorm.DB
}

func NewWriterDB(db *gorm.DB) *WriterDB {
	return &WriterDB{db: db}
}
