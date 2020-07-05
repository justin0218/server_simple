package models

import (
	"server_simple/api"
	"server_simple/internal/models/protoCompiles"
)

var BlogRoom blogRoomer

type blogRoom struct {
	tableName string
}

type blogRoomer interface {
	GetRoomList() (out protoCompiles.BlogRooms, err error)
}

func init() {
	BlogRoom = &blogRoom{
		tableName: "blog_rooms",
	}
}

func (this *blogRoom) GetRoomList() (out protoCompiles.BlogRooms, err error) {
	err = api.Mysql.Get().Find(&out.BlogRoomList).Error
	return
}
