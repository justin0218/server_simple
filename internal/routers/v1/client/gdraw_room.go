package client

import (
	"github.com/gin-gonic/gin"
	"server_simple/internal/services"
	"server_simple/pkg/resp"
)

var GdrawController gdrawController

func init() {
	GdrawController = &gdrawContron{}
}

type gdrawController interface {
	GetRoomList(c *gin.Context)
}

type gdrawContron struct {
}

func (ctr *gdrawContron) GetRoomList(c *gin.Context) {
	roomInfo := services.GdrawRoomer.GetRoomList()
	resp.RespOk(c, roomInfo)
	return
}
