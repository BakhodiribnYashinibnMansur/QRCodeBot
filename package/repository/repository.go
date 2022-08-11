package repository

import (
	"gorm.io/gorm"
)

type Reader interface {
}

type Writer interface {
}

type Repository struct {
	Reader
	Writer
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{Reader: NewReaderDB(db), Writer: NewWriterDB(db)}
}
