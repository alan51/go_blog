package models

import (
	"blog/helper"
	"github.com/astaxie/beego/orm"
)

type Zans struct {
	Id        int64
	ArticleId int64
	Ip        string `orm:"size(20)"`
	CreatedAt string `orm:"null"`
}

func GetZans() ([]orm.Params, error) {
	o := orm.NewOrm()
	zan := new(Zans)
	var zans []orm.Params
	_, err := o.QueryTable(zan).Filter("status", 1).Values(&zans)
	if err != nil {
		return zans, err
	}
	return zans, nil
}

//判断
func GetZansByIpAndArticleId(ip string, article_id int64) (v *Zans, err error) {
	o := orm.NewOrm()
	v = &Zans{Ip: ip, ArticleId: article_id}
	if err = o.QueryTable(new(Zans)).Filter("Ip", ip).Filter("ArticleId", article_id).RelatedSel().One(v); err == nil {
		return v, nil
	}
	return nil, err
}

//赞
func Zan(ip string, article_id int64) (err error) {
	o := orm.NewOrm()
	zan := new(Zans)
	zan.ArticleId = article_id
	zan.Ip = ip
	zan.CreatedAt = helper.NowFormat()
	_, err = o.Insert(zan)
	return
}
