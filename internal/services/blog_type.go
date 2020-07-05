package services

import (
	"server_simple/internal/models"
	"server_simple/internal/models/protoCompiles"
	"sync"
)

var BlogType blogtyper

func init() {
	BlogType = &blogtype{}
}

type blogtype struct {
	sync.Once
}

type blogtyper interface {
	GetBlogTypes(in protoCompiles.BlogType) (out protoCompiles.BlogTypes, err error)
}

func (this *blogtype) GetBlogTypes(in protoCompiles.BlogType) (out protoCompiles.BlogTypes, err error) {
	return models.BlogType.GetBlogTypes(in)
}
