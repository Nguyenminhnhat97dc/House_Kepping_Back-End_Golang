package database

/* import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db = *&gorm.DB{}
var dsn = "root:@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"

func DBConn() (db *gorm.DB) {
	if db == nil {
		dbc, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			panic("failed to connect database.")
		} else {
			fmt.Println("connect Successfull.")
		}
		db = dbc
	}
	return db
} */
