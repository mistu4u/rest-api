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
	
	ldb,err:=db.NewDBConn()
	if err!=nil{
		fmt.Println("error occurred")
	}
	return HiRepo{db: ldb.Gdb}
}

func (r HiRepo) SayHi() dto.MyMessage {
	var mes dto.MyMessage
	r.db.Table("message").First(&mes,"id=?","11e41e52-0333-47c2-af3b-791a68ba6ad0")
	return mes
}

var MessageRepoSet = wire.NewSet(NewRepo, wire.Bind(new(IHiRepo), new(HiRepo)))