package repo

import (
	"rest-api/dto"

	"github.com/google/wire"
)

type IHiRepo interface{
	SayHi() dto.MyMessage
}

type HiRepo struct {

}

func NewRepo() HiRepo {
	return HiRepo{}
}

func (r HiRepo) SayHi() dto.MyMessage {
	return dto.MyMessage{Message:"hi from repo"}
}

var MessageRepoSet = wire.NewSet(NewRepo, wire.Bind(new(IHiRepo), new(HiRepo)))