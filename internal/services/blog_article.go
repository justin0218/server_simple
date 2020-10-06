package services

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"server_simple/internal/models"
	"server_simple/internal/models/protoCompiles"
	"sync"
)

var Blog bloger

func init() {
	Blog = &blog{}
}

type blog struct {
	sync.Once
}

type bloger interface {
	GetBlogList(in protoCompiles.BlogArticle) (out protoCompiles.BlogList, err error)
	GetBlogDetail(c *gin.Context, in protoCompiles.BlogArticle) (protoCompiles.BlogDetail, error)
	GetRanking(limitNum int) (out protoCompiles.BlogList, err error)
	GetRecommend(limitNum int) (out protoCompiles.BlogList, err error)
	MakeGood(blogId int) (err error)
	Exist(blogId int) bool
}

func (this *blog) GetBlogList(in protoCompiles.BlogArticle) (out protoCompiles.BlogList, err error) {
	return models.Blog.GetBlogList(in)
}

func (this *blog) UpdateBlogView(in protoCompiles.BlogArticle) (err error) {
	return models.Blog.UpdateBlogView(in)
}

func (this *blog) GetBlogDetail(c *gin.Context, in protoCompiles.BlogArticle) (protoCompiles.BlogDetail, error) {
	res := protoCompiles.BlogDetail{}
	blogDetail, err := models.Blog.GetBlogDetail(in)
	if err != nil {
		return res, err
	}
	res.CurrentArticle = &blogDetail
	nextBlog, err := models.Blog.GetNextBlog(in)
	if err != nil {
		return res, err
	}
	res.NextArticle = &nextBlog
	prevBlog, err := models.Blog.GetPrevBlog(in)
	if err != nil {
		return res, err
	}
	res.PrevArticle = &prevBlog
	if res.CurrentArticle.Id != 0 {
		_ = this.UpdateBlogView(in)
	}
	return res, nil
}

func (this *blog) GetRanking(limitNum int) (out protoCompiles.BlogList, err error) {
	return models.Blog.GetRanking(limitNum)
}

func (this *blog) GetRecommend(limitNum int) (out protoCompiles.BlogList, err error) {
	return models.Blog.GetRecommend(limitNum)
}

func (this *blog) MakeGood(blogId int) (err error) {
	if !this.Exist(blogId) {
		err = fmt.Errorf("博客不存在")
		return
	}
	return models.Blog.MakeGood(blogId)
}

func (this *blog) Exist(blogId int) bool {
	if blogId == 0 {
		return true
	}
	blogInfo, _ := models.Blog.GetBlogDetail(protoCompiles.BlogArticle{Id: int32(blogId)})
	return blogInfo.Id != 0
}

func (this *blog) GetRoomList(limitNum int) (out protoCompiles.BlogList, err error) {
	return models.Blog.GetRanking(limitNum)
}
