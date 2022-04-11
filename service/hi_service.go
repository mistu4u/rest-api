package service

import (
	"rest-api/dto"
	"rest-api/repo"

	"github.com/google/wire"
)

type IHiService interface {
	SayHi() dto.MyMessage
}
type HiService struct {
	Repo repo.IHiRepo
}

func NewService(r repo.HiRepo) HiService {
	return HiService{Repo: r}
}

func (s HiService) SayHi() dto.MyMessage {
	return s.Repo.SayHi()
}

var HiServiceSet = wire.NewSet(NewService, wire.Bind(new(IHiService), new(HiService)))
