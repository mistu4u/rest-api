package repo

import (
	"fmt"
	"rest-api/db"
	"rest-api/dto"

	"github.com/google/wire"
	"gorm.io/gorm"
)

type IHiRepo interface{
	SayHi() dto.MyMessage
}

type HiRepo struct {
	db *gorm.DB
}

func NewRepo() HiRepo {
	ldb,err:=db.NewRepositories()
	if err!=nil{
		fmt.Println("error occurred")
	}
	k:=*ldb
	return HiRepo{db: k.Gdb}
}

func (r HiRepo) SayHi() dto.MyMessage {
	return dto.MyMessage{Message:"hi from repo"}
}

var MessageRepoSet = wire.NewSet(NewRepo, wire.Bind(new(IHiRepo), new(HiRepo)))