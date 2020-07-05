package services

import (
	"github.com/gogo/protobuf/proto"
	"github.com/gorilla/websocket"
	"server_simple/internal/models"
	"server_simple/internal/models/protoCompiles"
	"server_simple/pkg/tools"
	"strings"
	"time"
)

type ws struct {
	Conn *websocket.Conn
	Room string
}

func NewWs(conn *websocket.Conn, room string) wser {
	return &ws{
		Conn: conn,
		Room: room,
	}
}

type wser interface {
	Read()
}

func (this *ws) Read() {
	this.inroom()
	defer this.outroom()
	go this.pullFirstChatRecords(this.Room, 100)
	msgchan := make(chan []byte)
	for {
		go func() {
			select {
			case m, ok := <-msgchan:
				if !ok {
					return
				}
				wsBaseMsg := protoCompiles.WsMsgBase{}
				_ = proto.Unmarshal(m, &wsBaseMsg)
				switch wsBaseMsg.Event {
				case protoCompiles.Events_CHAT_CONTENT:
					chatContent := &protoCompiles.ChatContent{}
					_ = proto.Unmarshal(wsBaseMsg.Data, chatContent)
					tempCon := chatContent.Msg
					tempCon = strings.ReplaceAll(tempCon, " ", "")
					tempCon = strings.ReplaceAll(tempCon, "\t", "")
					tempCon = strings.ReplaceAll(tempCon, "\n", "")
					if tempCon == "" {
						return
					}
					chatContent.TimeString = time.Now().Format("2006-01-02 15:04:05")
					chatContent.Timer = int32(len([]byte(chatContent.Msg)))
					tools.Store.SetBlogChatRecordChan(protoCompiles.BlogChatRecord{
						Room:    this.Room,
						Content: chatContent.Msg,
					})
					wsBaseMsg.Data, _ = proto.Marshal(chatContent)
				}
				m, _ = proto.Marshal(&wsBaseMsg)
				this.broadcastToroom(m)
			}
		}()
		_, rmsg, err := this.Conn.ReadMessage()
		if err != nil {
			break
		}
		msgchan <- rmsg
	}
}

func (this *ws) broadcastToroom(msg []byte) {
	roomSumConn := tools.Store.GetRoomSumConn()
	conns := roomSumConn[this.Room]
	for _, val := range conns {
		_ = val.WriteMessage(websocket.BinaryMessage, msg)
	}
}

func (this *ws) inroom() {
	roomSumConn := tools.Store.GetRoomSumConn()
	conns := roomSumConn[this.Room]
	conns = append(conns, this.Conn)
	roomSumConn[this.Room] = conns
	tools.Store.SetRoomSumConn(roomSumConn)
}

func (this *ws) outroom() {
	roomSumConn := tools.Store.GetRoomSumConn()
	_ = this.Conn.Close()
	conns := roomSumConn[this.Room]
	for i, c := range conns {
		if c == this.Conn {
			conns = append(conns[:i], conns[i+1:]...)
		}
	}
	roomSumConn[this.Room] = conns
	tools.Store.SetRoomSumConn(roomSumConn)
}

func (this *ws) pullFirstChatRecords(room string, limit int) {
	data, _ := models.BlogChatRecord.GetMultiple(room, limit)
	wsBaseMsg := protoCompiles.WsMsgBase{}
	wsBaseMsg.Event = protoCompiles.Events_CHAT_RECORDS
	wsBaseMsg.Data, _ = proto.Marshal(&data)
	m, _ := proto.Marshal(&wsBaseMsg)
	_ = this.Conn.WriteMessage(websocket.BinaryMessage, m)
}
