package services

import (
	"server_simple/internal/models"
	"server_simple/internal/models/protoCompiles"
	"server_simple/pkg/tools"
)

var BlogRoom blogRoomer

type blogRoom struct {
}

type blogRoomer interface {
	GetRoomList() (out protoCompiles.BlogRooms, err error)
}

func init() {
	BlogRoom = &blogRoom{}
}

func (this *blogRoom) GetRoomList() (protoCompiles.BlogRooms, error) {
	out, err := models.BlogRoom.GetRoomList()
	for dx, val := range out.BlogRoomList {
		out.BlogRoomList[dx].Olnum = int32(len(tools.Store.GetRoomSumConn()[val.Id]))
	}
	return out, err
}
