package repo
import (
	"testing"
	
)
func TestHiRepo(t *testing.T){
	h := NewRepo()
	got:=h.SayHi()
	want:="Hi from Repo"
	if got!=want {
        t.Errorf("got %q, wanted %q", got, want)
    }
}