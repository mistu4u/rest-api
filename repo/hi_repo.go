package repo

import (
	"fmt"
	"github.com/google/wire"
	"gorm.io/gorm"
	"rest-api/db"
	"rest-api/dto"
	"rest-api/logger"
)

type IHiRepo interface {
	SayHi() dto.MyMessage
}

type HiRepo struct {
	db *gorm.DB
}

func NewRepo() HiRepo {
	ldb, err := db.NewDBConn()
	if err != nil {
		fmt.Println("error occurred in obtaining new db connection")
	}
	return HiRepo{db: ldb.Gdb}
}

func (r HiRepo) SayHi() dto.MyMessage {
	log := logger.NewLogger()
	log.Info("hello")
	var mes dto.MyMessage
	r.db.Table("message").First(&mes, "id=?", "11e41e52-0333-47c2-af3b-791a68ba6ad0")
	return mes
}

var MessageRepoSet = wire.NewSet(NewRepo, wire.Bind(new(IHiRepo), new(HiRepo)))
