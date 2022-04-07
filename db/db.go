package db

import (
	"fmt"

	"gorm.io/driver/postgres"
	_ "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Repositories struct {
	Gdb   *gorm.DB
}

//returns a new db connection
func NewRepositories() (*Repositories,error){
	dsn := "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable"
	ldb,err:=gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("error occurred")
		fmt.Errorf("error occurred in db connection")
		return nil, err
	}

	return &Repositories{Gdb:ldb},nil
}