package models

import (
	"github.com/gin-gonic/gin"
	"github.com/micro/protobuf/proto"
	"server_simple/internal/models/protoCompiles"
)

type errorer interface {
	JsonOK(payload proto.Message)
	JsonError(errcode protoCompiles.ErrorCodes, errmsg string)
}

type err struct {
	c *gin.Context
}

func (this *err) JsonOK(payload proto.Message) {
	resData, _ := proto.Marshal(payload)
	this.c.Writer.Write(resData)
	return
}

func (this *err) JsonError(errcode protoCompiles.ErrorCodes, errmsg string) {
	res := &protoCompiles.Resp{}
	res.Code = errcode
	res.Msg = errmsg
	resData, _ := proto.Marshal(res)
	this.c.Writer.Write(resData)
	return
}

func SetGinContext(c *gin.Context) errorer {
	e := &err{}
	e.c = c
	return e
}
