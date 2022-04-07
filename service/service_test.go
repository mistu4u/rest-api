package service

import (
	"fmt"
	"reflect"
	"rest-api/dto"
	"rest-api/mocks"
	"testing"

	"github.com/go-playground/assert/v2"
)

//use mock to test the functionality
func TestUsingMock(t *testing.T){
	r:=dto.MyMessage{Message: "mocking"}
	mRepo:=new(mocks.IHiRepo)
	mRepo.On("SayHi").Return(r)
	ns:=HiService{Repo: mRepo}
	m:=ns.SayHi()
	fmt.Println(reflect.TypeOf(m))
    assert.Equal(t,m,dto.MyMessage{Message: "mocking"})
}