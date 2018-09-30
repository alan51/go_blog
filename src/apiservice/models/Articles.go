package models

import (
	"github.com/astaxie/beego/orm"
)

type Articles struct {
	Id          int64
	Title       string `orm:"size(255)" valid:"Required;MaxSize:255;MinSize:2"`
	Desc        string `orm:"size(1024)"`
	Content     string `orm:"type(text)"`
	IndexImg    string `orm:"size(150)"`
	CommentNum  int    `orm:"default(0)"`
	ZanNum      int    `orm:"default(0)"`
	ViewNum     int    `orm:"default(0)"`
	IsRecommend int16  `orm:"default(1)"`
	IsTop       int16  `orm:"default(2)"`
	Status      int16  `orm:"default(1)"`
	Tags        string
	CreatedAt   string `orm:"null"`
	UpdatedAt   string
	//User        *Users      `orm:"rel(fk)"`
	Category    *Categories `orm:"rel(fk)"`
	Keywords    string      `orm:"size(255);null"`
}

func GetAllArticles() ([]orm.Params, error) {
	o := orm.NewOrm()
	article := new(Articles)
	var articles []orm.Params
	_, err := o.QueryTable(article).Filter("status", 1).Values(&articles)
	if err != nil {
		return articles, err
	}
	return articles, nil
}