package models

import (
	"server_simple/api"
	"server_simple/internal/models/protoCompiles"
)

var BlogChatRecord blogChatRecorder

func init() {
	BlogChatRecord = &blogChatRecord{}
}

type blogChatRecord struct {
}

type blogChatRecorder interface {
	Create(in protoCompiles.BlogChatRecord) (out protoCompiles.BlogChatRecord, err error)
	GetMultiple(room string, limit int) (out protoCompiles.BlogChatRecords, err error)
}

func (this *blogChatRecord) Create(in protoCompiles.BlogChatRecord) (out protoCompiles.BlogChatRecord, err error) {
	err = api.Mysql.Get().Omit("xxx_unrecognized", "xxx_sizecache", "create_time", "update_time").Create(&in).Error
	out = in
	return
}

func (this *blogChatRecord) GetMultiple(room string, limit int) (out protoCompiles.BlogChatRecords, err error) {
	err = api.Mysql.Get().Where("room = ?", room).Order("id desc").Find(&out.BlogChatRecordList).Limit(limit).Error
	return
}
