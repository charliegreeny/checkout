package config

import (
	"database/sql"
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDb() (*gorm.DB, error) {
	addr := os.Getenv("DB_ADDRESS")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/checkout?charset=utf8mb4&parseTime=True&loc=Local", user, password, addr)
	sqlDB, err := sql.Open("mysql", dsn) 
	if err != nil {
		return nil, err
	}
	return gorm.Open(mysql.New(mysql.Config{Conn: sqlDB,}), &gorm.Config{})
}
