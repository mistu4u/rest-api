package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http/httptest"
	"rest-api/dto"
	"rest-api/mocks"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
)

func TestApi(t *testing.T) {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	mService := mocks.IHiService{}
	r := dto.MyMessage{Message: "mocking"}
	mService.On("SayHi").Return(r)
	hApi := HiAPI{HiService: &mService}
	hApi.SayHi(c)
	assert.Equal(t, 200, w.Code)
	b, _ := ioutil.ReadAll(w.Body)
	var testMessage dto.MyMessage
	fmt.Println(string(b))
	json.Unmarshal(b, &testMessage)
	assert.Equal(t, r, testMessage)
}
