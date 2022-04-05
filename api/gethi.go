package api

import (
	"net/http"
	"rest-api/service"

	"github.com/gin-gonic/gin"
)

type IApi interface {
	SayHi(c *gin.Context)
}
type HiAPI struct {
	HiService service.HiService
}

func NewApi(h service.HiService) HiAPI {
	return HiAPI{HiService: h}
}

func (h *HiAPI) SayHi(c *gin.Context) {
	m := h.HiService.SayHi()
	c.JSON(http.StatusOK, gin.H{"message": m})
}
