package ws

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"server_simple/internal/services"
	"strconv"
	"sync"
	//"server_simple/internal/models"
)

var Ws wser

type ws struct {
}

func init() {
	Ws = &ws{}
}

type wser interface {
	ChatRoom(c *gin.Context)
	Gdraw(c *gin.Context)
}

func (this *ws) ChatRoom(c *gin.Context) {
	room := c.Query("room")
	wsconnVal, _ := c.Get("wsconn")
	if wsconn, ok := wsconnVal.(*websocket.Conn); ok {
		wsClient := services.NewWs(wsconn, room)
		go wsClient.Read()
	}

}

type Pool struct {
	Queue chan int
	Wg    *sync.WaitGroup
}

func NewPool(cap, total int) *Pool {
	if cap < 1 {
		cap = 1
	}
	p := &Pool{
		Queue: make(chan int, cap),
		Wg:    new(sync.WaitGroup),
	}
	p.Wg.Add(total)
	return p
}

func (p *Pool) AddOne() {
	p.Queue <- 1
}

func (p *Pool) DelOne() {
	<-p.Queue
	p.Wg.Done()
}

func (this *ws) Gdraw(c *gin.Context) {
	room := c.Query("room_id")
	wsconnVal, _ := c.Get("wsconn")
	if wsconn, ok := wsconnVal.(*websocket.Conn); ok {
		uidstr := c.Query("uid")
		token := c.Query("token")
		var (
			uidInt64 int64
		)
		uidInt64, _ = strconv.ParseInt(uidstr, 10, 64)
		wsClient := services.NewGdrawWs(wsconn, room, uidInt64, token)
		go wsClient.Read()
		go wsClient.Write()
	}
}
