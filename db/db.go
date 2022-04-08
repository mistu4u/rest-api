package db

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DBConn struct {
	Gdb   *gorm.DB
}

//returns a new db connection
func NewDBConn() (*DBConn,error){
	dsn := "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable"
	ldb,err:=gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info)})
	if err != nil {
		fmt.Println("error occurred in db connection")
		return nil, err
	}

	return &DBConn{Gdb:ldb},nil
}