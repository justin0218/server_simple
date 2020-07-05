package models

import (
	"server_simple/api"
	"server_simple/internal/models/protoCompiles"
)

var BlogComment blogCommenter

func init() {
	BlogComment = &blogComment{
		tableName: "blog_comments",
	}
}

type blogComment struct {
	tableName string
}

type blogCommenter interface {
	GetCommentList(blogId int) (out protoCompiles.BlogComments, err error)
	CreateComment(in protoCompiles.BlogComment) (err error)
}

func (this *blogComment) GetCommentList(blogId int) (out protoCompiles.BlogComments, err error) {
	err = api.Mysql.Get().Table(this.tableName).Where("blog_id = ?", blogId).Limit(100).Order("id desc").Find(&out.BlogCommentList).Error
	err = api.Mysql.Get().Table(this.tableName).Where("blog_id = ?", blogId).Count(&out.Total).Error
	return
}

func (this *blogComment) CreateComment(in protoCompiles.BlogComment) (err error) {
	err = api.Mysql.Get().Table(this.tableName).Omit("xxx_unrecognized", "xxx_sizecache", "create_time", "update_time").Create(&in).Error
	return
}
