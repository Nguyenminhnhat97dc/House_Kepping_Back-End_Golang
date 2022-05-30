package connectdatabase

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//var db *gorm.DB
//var dsn = "root:@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Loca

var dsn = "sql6496052:JVUfiJ9mBJ@tcp(sql6.freemysqlhosting.net:3306)/sql6496052?charset=utf8mb4&parseTime=True&loc=Local"

func DBConn() (db *gorm.DB) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database.")
	} else {
		fmt.Println("connect Successfull.")
	}
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalln(err)
	}
	sqlDB.SetConnMaxLifetime(1 * time.Second)
	return db
}
