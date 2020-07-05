package job

import (
	"server_simple/api"
	"server_simple/internal/models"
	"server_simple/pkg/tools"
	"server_simple/pkg/tools/store"
	"time"
)

func NewJob() jober {
	return &job{}
}

type job struct {
}

type jober interface {
	BlogChatRecord()
	GdrawData()
}

func (this *job) BlogChatRecord() {
	for {
		data := tools.Store.GetBlogChatRecordChan()
		_, _ = models.BlogChatRecord.Create(data)
		time.Sleep(time.Millisecond * 100)
	}
}

func (this *job) GdrawData() {
	for {
		select {
		case msg, ok := <-store.GdrawDataChan:
			if !ok {
				return
			}
			err := models.GdrawUserAnswerer.Create(msg)
			if err != nil {
				api.Log.Get().Error("job GdrawData GdrawUserAnswerer Create err:%v", err)
			}
		}
	}
}
