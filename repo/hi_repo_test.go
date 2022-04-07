package repo

import (
	"rest-api/dto"
	"testing"
)
func TestHiRepo(t *testing.T){
	h := NewRepo()
	got:= h.SayHi()
	want:= dto.MyMessage{Message:"hi from repo" }
	if got!=want {
        t.Errorf("got %q, wanted %q", got, want)
    }
}

