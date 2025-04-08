package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func Connect() {
	dsn := "host=localhost user=postgres password=20021512 dbname=dessert_shop port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("数据库连接失败:", err)
	}
	DB = db
	fmt.Println("数据库连接成功")
}
