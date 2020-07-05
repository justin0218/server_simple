package services

import (
	"github.com/parnurzeal/gorequest"
	"server_simple/internal/cache"
	"server_simple/internal/models/protoCompiles"
)

var PublicRead publicReader

type publicReader interface {
	ReadNetFile(netUrl string) protoCompiles.FileContent
}

type publicRead struct {
}

func init() {
	PublicRead = &publicRead{}
}

func (this *publicRead) ReadNetFile(netUrl string) protoCompiles.FileContent {
	res := cache.PublicRead.ReadFile(netUrl)
	if res.Txt == "" {
		gorequest.New().Get(netUrl).EndBytes(func(response gorequest.Response, body []byte, errs []error) {
			_ = cache.PublicRead.SetFile(netUrl, body)
			res.Txt = string(body)
		})
		return res
	}
	return res
}
