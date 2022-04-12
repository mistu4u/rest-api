package api

import (
	"net/http"
	"rest-api/service"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

type IApi interface {
	SayHi(c *gin.Context)
}

type HiAPI struct {
	HiService service.IHiService
}

func NewApi(h service.HiService) HiAPI {
	return HiAPI{HiService: h}
}

func (h HiAPI) SayHi(c *gin.Context) {
	m := h.HiService.SayHi()
	c.JSON(http.StatusOK, m)
}

var HiAPISet = wire.NewSet(NewApi, wire.Bind(new(IApi), new(HiAPI)))
