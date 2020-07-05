package services

import (
	"server_simple/internal/models"
	"server_simple/pkg/tools/store"
	"strconv"
)

var GdrawRoomer gdrawRoomer

func init() {
	GdrawRoomer = &gdrawRoom{}
}

type gdrawRoom struct {
}

type gdrawRoomer interface {
	GetRoomList() (res GetRoomListRes)
	JoinRoom(uid int64, roomId string) (err error)
}

type GetRoomListRes struct {
	List []GetRoomListItem `json:"list"`
}

type GetRoomListItem struct {
	Uinfo    models.GdrawUser       `json:"uinfo"`
	RoomInfo EVENT_GAME_STATUS_DATA `json:"room_info"`
}

func (this *gdrawRoom) GetRoomList() (res GetRoomListRes) { //肯定房主才会来创建房间
	rooms := []int64{}
	store.ROOM_READY_STATUS.Range(func(key, value interface{}) bool {
		roomUidInt, _ := strconv.Atoi(key.(string))
		rooms = append(rooms, int64(roomUidInt))
		roomLeaderInfos, _ := models.GdrawUserer.FindGdrawUserWithIds(rooms)
		for _, val := range roomLeaderInfos {
			item := GetRoomListItem{}
			item.Uinfo = val
			item.RoomInfo = value.(EVENT_GAME_STATUS_DATA)
			res.List = append(res.List, item)
		}
		return true
	})
	return
}

func (this *gdrawRoom) JoinRoom(uid int64, roomId string) (err error) {
	var uids []int64
	if v, ok := store.ROOM_ALL_PERSON.Load(roomId); ok {
		uids = v.([]int64)
		uids = append(uids, uid)
		store.ROOM_ALL_PERSON.Store(roomId, uids)
		return
	}
	uids = append(uids, uid)
	store.ROOM_ALL_PERSON.Store(roomId, uids)
	return
}
