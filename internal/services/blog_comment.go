package services

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"server_simple/api"
	"server_simple/internal/models"
	"server_simple/internal/models/protoCompiles"
	"time"
)

var BlogComment blogCommenter

func init() {
	BlogComment = &blogComment{}
}

type blogComment struct {
}

type blogCommenter interface {
	GetCommentList(blogId int) (out protoCompiles.BlogComments, err error)
	SubmitComment(c *gin.Context, in protoCompiles.BlogComment) (err error)
}

func (this *blogComment) GetCommentList(blogId int) (out protoCompiles.BlogComments, err error) {
	return models.BlogComment.GetCommentList(blogId)
}

func (this *blogComment) SubmitComment(c *gin.Context, in protoCompiles.BlogComment) (err error) {
	if !Blog.Exist(int(in.Id)) {
		err = fmt.Errorf("博客不存在")
		return
	}
	ip := c.ClientIP()
	key := fmt.Sprintf("%s:%s:%d", time.Now().Format("2006-01-02"), ip, in.BlogId)
	num, err := api.Rds.Get().Incr(key).Result()
	api.Rds.Get().Expire(key, time.Hour*24)
	if err != nil {
		return
	}
	if num > 10 {
		err = fmt.Errorf("已超过今日评论最大上限")
		return
	}
	tempUserName := uuid.New().String()
	in.Name = tempUserName
	return models.BlogComment.CreateComment(in)
}
