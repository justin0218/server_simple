package models

import (
	"server_simple/api"
	"server_simple/internal/models/protoCompiles"
)

var BlogType blogtyper

func init() {
	BlogType = &blogtype{
		tableName: "blog_types",
	}
}

type blogtype struct {
	tableName string
}

type blogtyper interface {
	GetBlogTypes(in protoCompiles.BlogType) (out protoCompiles.BlogTypes, err error)
}

func (this *blogtype) GetBlogTypes(in protoCompiles.BlogType) (out protoCompiles.BlogTypes, err error) {
	err = api.Mysql.Get().Table(this.tableName).Find(&out.BlogTypeList).Error
	return
}
