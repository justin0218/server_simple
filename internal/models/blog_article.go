package models

import (
	"github.com/jinzhu/gorm"
	"server_simple/api"
	"server_simple/internal/models/protoCompiles"
)

var Blog bloger

func init() {
	Blog = &blog{
		tableName: "blog_articles",
	}
}

type blog struct {
	tableName string
}

type bloger interface {
	GetBlogList(in protoCompiles.BlogArticle) (out protoCompiles.BlogList, err error)
	GetBlogDetail(in protoCompiles.BlogArticle) (out protoCompiles.BlogArticle, err error)
	GetNextBlog(in protoCompiles.BlogArticle) (out protoCompiles.BlogArticle, err error)
	GetPrevBlog(in protoCompiles.BlogArticle) (out protoCompiles.BlogArticle, err error)
	UpdateBlogView(in protoCompiles.BlogArticle) (err error)
	GetRanking(limitNum int) (out protoCompiles.BlogList, err error)
	GetRecommend(limitNum int) (out protoCompiles.BlogList, err error)
	MakeGood(blogId int) (err error)
}

func (this *blog) GetBlogList(in protoCompiles.BlogArticle) (out protoCompiles.BlogList, err error) {
	result := api.Mysql.Get().Table(this.tableName)
	if in.Name != "" {
		result = result.Where("name LIKE ?", "%"+in.Name+"%")
	}
	if in.Type != -1 {
		result = result.Where("type = ?", in.Type)
	}
	err = result.Where("id != ?", 35).Order("recommended desc").Find(&out.BlogArticleList).Error
	return
}

func (this *blog) GetBlogDetail(in protoCompiles.BlogArticle) (out protoCompiles.BlogArticle, err error) {
	result := api.Mysql.Get().Table(this.tableName)
	err = result.Where("id = ?", in.Id).Find(&out).Error
	return
}

func (this *blog) GetNextBlog(in protoCompiles.BlogArticle) (out protoCompiles.BlogArticle, err error) {
	result := api.Mysql.Get().Table(this.tableName)
	err = result.Where("id > ?", in.Id).Where("id != ?", 35).First(&out).Error
	return
}

func (this *blog) GetPrevBlog(in protoCompiles.BlogArticle) (out protoCompiles.BlogArticle, err error) {
	result := api.Mysql.Get().Table(this.tableName)
	err = result.Where("id < ?", in.Id).Where("id != ?", 35).Order("id desc").First(&out).Error
	return
}

func (this *blog) UpdateBlogView(in protoCompiles.BlogArticle) (err error) {
	result := api.Mysql.Get().Table(this.tableName)
	err = result.Where("id = ?", in.Id).UpdateColumn("view", gorm.Expr("view + ?", 1)).Error
	return
}

func (this *blog) GetRanking(limitNum int) (out protoCompiles.BlogList, err error) {
	result := api.Mysql.Get().Table(this.tableName)
	err = result.Where("id != ?", 35).Order("view desc").Limit(limitNum).Find(&out.BlogArticleList).Error
	return
}

func (this *blog) GetRecommend(limitNum int) (out protoCompiles.BlogList, err error) {
	result := api.Mysql.Get().Table(this.tableName)
	err = result.Where("id != ?", 35).Where("recommended = ?", 1).Order("id desc").Limit(limitNum).Find(&out.BlogArticleList).Error
	return
}

func (this *blog) MakeGood(blogId int) (err error) {
	result := api.Mysql.Get().Table(this.tableName)
	err = result.Where("id = ?", blogId).UpdateColumn("good_num", gorm.Expr("good_num + ?", 1)).Error
	return
}
