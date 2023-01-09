package config

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func Connect() {
	// Create database "GoLang-test" using myphpadmin and edit in this "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	d, err := gorm.Open("mysql", "root:@tcp(127.0.0.1:3306)/GoLang-test?charset=utf8mb4&parseTime=True&loc=Local")
	fmt.Println("Connected with 'GoLang-test'")
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
