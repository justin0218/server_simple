package models

import (
	"server_simple/api"
	"server_simple/internal/models/protoCompiles"
)

var BlogView blogviewer

func init() {
	BlogView = &blogview{
		tableName: "blog_views",
	}
}

type blogview struct {
	tableName string
}

type blogviewer interface {
	CreateBlogView(in protoCompiles.BlogView) (out protoCompiles.BlogView, err error)
}

func (this *blogview) CreateBlogView(in protoCompiles.BlogView) (out protoCompiles.BlogView, err error) {
	intoReq := protoCompiles.BlogView{}
	intoReq.Country = in.Country
	intoReq.Region = in.Region
	intoReq.City = in.City
	intoReq.Isp = in.Isp
	intoReq.BlogId = in.BlogId
	intoReq.Ip = in.Ip
	err = api.Mysql.Get().Table(this.tableName).Omit("xxx_unrecognized", "xxx_sizecache", "create_time", "update_time").Create(&intoReq).Error
	return
}
