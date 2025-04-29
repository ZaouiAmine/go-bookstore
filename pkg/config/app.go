package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	Db *gorm.DB
)

// ConnectDB function to connect to the database
func ConnectDB() {
	// Connect to the database
	var err error
	D, err := gorm.Open("mysql", "root:zaoui111@tcp(localhost:3306)/bookstore?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	// Migrate the schema
	Db = D
}

func GetDB() *gorm.DB {
	return Db
}

// CloseDB function to close the database connection
func CloseDB() {
	if Db != nil {
		Db.Close()
	}
}
