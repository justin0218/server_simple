package client

import (
	"github.com/gin-gonic/gin"
	"github.com/gogo/protobuf/proto"
	"io/ioutil"
	"server_simple/api"
	"server_simple/internal/models"
	"server_simple/internal/models/protoCompiles"
	"server_simple/internal/services"
	"strconv"
	"strings"
)

var BlogController blogController

func init() {
	BlogController = &blogContron{}
}

type blogController interface {
	GetList(c *gin.Context)
	GetDetail(c *gin.Context)
	GetTypes(c *gin.Context)
	GetRanking(c *gin.Context)
	GetRecommend(c *gin.Context)
	GetCommentList(c *gin.Context)
	SubmitComment(c *gin.Context)
	MakeGood(c *gin.Context)
	GetRoomList(c *gin.Context)
}

type blogContron struct {
}

func (ctr *blogContron) GetList(c *gin.Context) {
	es := models.SetGinContext(c)
	tp := c.DefaultQuery("tp", "-1")
	tpInt, err := strconv.Atoi(tp)
	if err != nil {
		tpInt = -1
	}
	name := c.Query("name")
	req := protoCompiles.BlogArticle{}
	req.Name = name
	req.Type = int64(tpInt)
	res, err := services.Blog.GetBlogList(req)
	if err != nil {
		api.Log.Get().Error("blogContron GetList err:%v", err)
	}
	es.JsonOK(&res)
	return
}

func (ctr *blogContron) GetDetail(c *gin.Context) {
	es := models.SetGinContext(c)
	id := c.Query("id")
	idInt, _ := strconv.Atoi(id)
	req := protoCompiles.BlogArticle{}
	req.Id = int32(idInt)
	res, err := services.Blog.GetBlogDetail(c, req)
	if err != nil {
		api.Log.Get().Error("blogContron GetDetail err:%v", err)
	}
	es.JsonOK(&res)
	return
}

func (ctr *blogContron) GetTypes(c *gin.Context) {
	es := models.SetGinContext(c)
	req := protoCompiles.BlogType{}
	res, err := services.BlogType.GetBlogTypes(req)
	if err != nil {
		api.Log.Get().Error("blogContron GetList err:%v", err)
	}
	es.JsonOK(&res)
	return
}

func (ctr *blogContron) GetRanking(c *gin.Context) {
	es := models.SetGinContext(c)
	limit := c.Query("limit")
	limitNum, _ := strconv.Atoi(limit)
	if limitNum == 0 {
		limitNum = 1
	}
	res, _ := services.Blog.GetRanking(limitNum)
	es.JsonOK(&res)
	return
}

func (ctr *blogContron) GetRecommend(c *gin.Context) {
	es := models.SetGinContext(c)
	limit := c.Query("limit")
	limitNum, _ := strconv.Atoi(limit)
	if limitNum == 0 {
		limitNum = 1
	}
	res, _ := services.Blog.GetRecommend(limitNum)
	es.JsonOK(&res)
	return
}

func (ctr *blogContron) GetCommentList(c *gin.Context) {
	es := models.SetGinContext(c)
	blogId := c.Query("blog_id")
	blogIdNum, _ := strconv.Atoi(blogId)
	res, _ := services.BlogComment.GetCommentList(blogIdNum)
	es.JsonOK(&res)
	return
}

func (ctr *blogContron) SubmitComment(c *gin.Context) {
	es := models.SetGinContext(c)
	bodyBytes, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		es.JsonError(protoCompiles.ErrorCodes_INVALID_PARAMS, err.Error())
		return
	}
	defer c.Request.Body.Close()
	req := protoCompiles.BlogComment{}
	err = proto.Unmarshal(bodyBytes, &req)
	if err != nil {
		es.JsonError(protoCompiles.ErrorCodes_INVALID_PARAMS, err.Error())
		return
	}
	tempCon := req.Content
	tempCon = strings.ReplaceAll(tempCon, " ", "")
	tempCon = strings.ReplaceAll(tempCon, "\t", "")
	tempCon = strings.ReplaceAll(tempCon, "\n", "")
	if tempCon == "" {
		es.JsonError(protoCompiles.ErrorCodes_INVALID_PARAMS, "评论内容不能为空")
		return
	}
	err = services.BlogComment.SubmitComment(c, req)
	if err != nil {
		es.JsonError(protoCompiles.ErrorCodes_INVALID_PARAMS, err.Error())
		return
	}
	es.JsonOK(nil)
	return
}

func (ctr *blogContron) MakeGood(c *gin.Context) {
	es := models.SetGinContext(c)
	blogId := c.Query("blog_id")
	blogIdNum, _ := strconv.Atoi(blogId)
	err := services.Blog.MakeGood(blogIdNum)
	if err != nil {
		es.JsonError(protoCompiles.ErrorCodes_INVALID_PARAMS, err.Error())
		return
	}
	es.JsonOK(nil)
	return
}

func (ctr *blogContron) GetRoomList(c *gin.Context) {
	es := models.SetGinContext(c)
	res, err := services.BlogRoom.GetRoomList()
	if err != nil {
		es.JsonError(protoCompiles.ErrorCodes_INVALID_PARAMS, err.Error())
		return
	}
	es.JsonOK(&res)
	return
}
