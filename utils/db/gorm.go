package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

// root@tcp(host.docker.internal:3306)/login?charset=utf8mb4&parseTime=True&loc=UTC
func GormMysql() *gorm.DB {
	dsn := os.Getenv("DB_OSN")
	fmt.Println("databse", dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("gorm.open", err)
	}
	return db

}
