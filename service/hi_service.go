package service

import "rest-api/repo"

type IHiService interface {
	SayHi() string
}
type HiService struct {
	Repo repo.HiRepo
}

func NewService(r repo.HiRepo) HiService {
	return HiService{Repo: r}
}

func (s *HiService) SayHi() string {
	return s.Repo.SayHi()
}
