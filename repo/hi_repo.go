package repo

type HiRepo struct {
}

func NewRepo() HiRepo {
	return HiRepo{}
}

func (r *HiRepo) SayHi() string {
	return "Hi from Repo"
}
