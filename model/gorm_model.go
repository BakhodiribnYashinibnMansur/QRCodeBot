package model

import "gorm.io/gorm"

type Account struct {
	gorm.Model
	ID string
	Username  string
	 
}
