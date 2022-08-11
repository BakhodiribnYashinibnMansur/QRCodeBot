package repository

import "gorm.io/gorm"

type ReaderDB struct {
	db *gorm.DB
}

func NewReaderDB(db *gorm.DB) *ReaderDB {
	return &ReaderDB{db: db}
}

// func (repo *ReaderDB) CreateUser(user model.User, logrus *logrus.Logger) (int, error)
