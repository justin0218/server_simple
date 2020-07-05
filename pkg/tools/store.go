package tools

import (
	"github.com/gorilla/websocket"
	"server_simple/internal/models/protoCompiles"
)

var Store storer

type RoomInfo struct {
	Uid  int64
	Conn *websocket.Conn
}

type store struct {
	roomSumConn        map[string][]*websocket.Conn
	blogChatRecordChan chan protoCompiles.BlogChatRecord
	roomSumInfo        map[string][]RoomInfo
}

type storer interface {
	GetRoomSumConn() map[string][]*websocket.Conn
	SetRoomSumConn(rs map[string][]*websocket.Conn)
	SetBlogChatRecordChan(in protoCompiles.BlogChatRecord)
	GetBlogChatRecordChan() protoCompiles.BlogChatRecord
	GetRoomSumInfo() map[string][]RoomInfo
	SetRoomSumInfo(rs map[string][]RoomInfo)
}

func init() {
	Store = &store{
		roomSumConn:        make(map[string][]*websocket.Conn),
		blogChatRecordChan: make(chan protoCompiles.BlogChatRecord),
		roomSumInfo:        make(map[string][]RoomInfo),
	}
}

func (this *store) GetRoomSumConn() map[string][]*websocket.Conn {
	return this.roomSumConn
}

func (this *store) SetRoomSumConn(rs map[string][]*websocket.Conn) {
	this.roomSumConn = rs
}

func (this *store) GetRoomSumInfo() map[string][]RoomInfo {
	return this.roomSumInfo
}

func (this *store) SetRoomSumInfo(rs map[string][]RoomInfo) {
	this.roomSumInfo = rs
}

func (this *store) SetBlogChatRecordChan(in protoCompiles.BlogChatRecord) {
	this.blogChatRecordChan <- in
}

func (this *store) GetBlogChatRecordChan() protoCompiles.BlogChatRecord {
	return <-this.blogChatRecordChan
}
