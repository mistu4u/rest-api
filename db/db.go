package db

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DBConn struct {
	Gdb *gorm.DB
}

//returns a new db connection
func NewDBConn() (*DBConn, error) {
	dbhost := os.Getenv("DB_HOST")
	dsn := fmt.Sprintf("host=%s user=postgres password=postgres dbname=postgres port=5432 sslmode=disable", dbhost)
	fmt.Println(dsn)
	ldb, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info)})
	if err != nil {
		fmt.Println("error occurred in db connection")
		return nil, err
	}

	return &DBConn{Gdb: ldb}, nil
}
