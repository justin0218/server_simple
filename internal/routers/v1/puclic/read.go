package puclic

import (
	"github.com/gin-gonic/gin"
	"server_simple/internal/models"
	"server_simple/internal/services"
)

var PublicController publicController

func init() {
	PublicController = &publicContron{}
}

type publicController interface {
	ReadNetFile(c *gin.Context)
}

type publicContron struct {
}

func (this *publicContron) ReadNetFile(c *gin.Context) {
	es := models.SetGinContext(c)
	k := c.Query("key")
	res := services.PublicRead.ReadNetFile(k)
	es.JsonOK(&res)
	return
}
