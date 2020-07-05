package services

import (
	"server_simple/internal/models"
	"server_simple/internal/models/protoCompiles"
	"sync"
)

var BlogView blogviewer

func init() {
	BlogView = &blogview{}
}

type blogview struct {
	sync.Once
}

type blogviewer interface {
	CreateBlogView(in protoCompiles.BlogView) (out protoCompiles.BlogView, err error)
}

func (this *blogview) CreateBlogView(in protoCompiles.BlogView) (out protoCompiles.BlogView, err error) {
	return models.BlogView.CreateBlogView(in)
}
