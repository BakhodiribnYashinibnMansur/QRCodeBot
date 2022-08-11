package repository

import (
	"telegrambot/model"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewSQLliteDB() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("database/telegram.db"), &gorm.Config{})
	if err != nil {
		logrus.Fatal("failed to connect database")
		return nil, err
	}
	account := &model.Account{}

	db.AutoMigrate(account)
	//Insert content
	//   db. Create (& product {Title: "new phone", code: "d42", price: 1000})
	//   db. Create (& product {Title: "new computer", code: "d43", price: 3500})

	//   //Read content
	//   var product Product
	//   db.First(&product, 1) // find product with integer primary key
	//   db.First(&product, "code = ?", "D42") // find product with code D42

	//   //Update operation: update a single field
	//   db.Model(&product).Update("Price", 2000)

	//   //Update operation: update multiple fields
	//   db.Model(&product).Updates(Product{Price: 2000, Code: "F42"}) // non-zero fields
	//   db.Model(&product).Updates(map[string]interface{}{"Price": 2000, "Code": "F42"})

	//   //Delete operation:
	//   db.Delete(&product, 1)
	return db, err
}
