package cache

import (
	"server_simple/api"
	"server_simple/internal/models/protoCompiles"
	"time"
)

var PublicRead publicReader

type publicReader interface {
	ReadFile(k string) protoCompiles.FileContent
	SetFile(k string, v interface{}) error
}

type publicRead struct {
}

func init() {
	PublicRead = &publicRead{}
}

func (this *publicRead) ReadFile(k string) protoCompiles.FileContent {
	r := protoCompiles.FileContent{}
	res, _ := api.Rds.Get().Get(k).Result()
	r.Txt = res
	return r
}

func (this *publicRead) SetFile(k string, v interface{}) error {
	return api.Rds.Get().Set(k, v, time.Hour*24*30).Err()
}
