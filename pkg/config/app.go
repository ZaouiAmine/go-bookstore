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
	D, err = gorm.Open("mysql", "root:root@simplerest?")
	if err != nil {
		panic("failed to connect to database")
	}
	Db = D
	