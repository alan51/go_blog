package models

import (
	"blog/helper"
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"

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
	User        *Users      `orm:"rel(fk)"`
	Category    *Categories `orm:"rel(fk)"`
	Keywords    string      `orm:"size(255);null"`
}

func GetArticles() ([]orm.Params, error) {
	o := orm.NewOrm()
	article := new(Articles)
	var articles []orm.Params
	_, err := o.QueryTable(article).Filter("status", 1).Values(&articles)
	if err != nil {
		return articles, err
	}
	return articles, nil
}

//  新增网站分类
func AddArticles(m *Articles) (id int64, err error) {
	o := orm.NewOrm()
	article := new(Articles)
	article.Title = m.Title
	//article.Category = Categories{Id:m.Ca}
	article.Desc = m.Desc
	article.Content = m.Content
	article.IndexImg = m.IndexImg
	article.IsRecommend = int16(m.IsRecommend)
	article.IsTop = int16(m.IsTop)
	article.Status = int16(m.Status)
	article.Tags = m.Tags
	article.CreatedAt = helper.NowFormat()
	article.UpdatedAt = article.CreatedAt
	article.Category = m.Category
	article.Keywords = m.Keywords
	article.User = m.User
	id, err = o.Insert(article)
	return
}

// 详情
func GetArticlesById(id int64) (v *Articles, err error) {
	o := orm.NewOrm()
	v = &Articles{Id: id}
	if err = o.QueryTable(new(Articles)).Filter("Id", id).RelatedSel().One(v); err == nil {
		return v, nil
	}
	return nil, err
}

// 上一篇
func GetPreArticlesById(id int64) (v *Articles, err error) {
	o := orm.NewOrm()
	v = &Articles{Id: id}
	if err = o.QueryTable(new(Articles)).Filter("Id__lt", id).One(v, "id", "title"); err == nil {
		return v, nil
	}
	return nil, err
}

//下一篇
func GetNextArticlesById(id int64) (v *Articles, err error) {
	o := orm.NewOrm()
	v = &Articles{Id: id}
	if err = o.QueryTable(new(Articles)).Filter("Id__gt", id).One(v, "id", "title"); err == nil {
		return v, nil
	}
	return nil, err
}

// 列表
func GetAllArticles(query map[string]string, fields []string, sortby []string, order []string,
	offset int, limit int) (ml []interface{}, count int, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Articles))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		qs = qs.Filter(k, v)
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, 0, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, 0, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, 0, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, 0, errors.New("Error: unused 'order' fields")
		}
	}

	var l []Articles
	qs = qs.OrderBy(sortFields...).RelatedSel()
	if _, err = qs.Limit(limit, offset).All(&l, "Id", "Title", "Category", "User", "Desc", "IndexImg", "IsRecommend", "IsTop", "Status", "Tags", "CreatedAt", "UpdatedAt", "Keywords", "ViewNum"); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		count, _ := qs.Count()
		return ml, int(count), nil
	}
	return nil, 0, err
}

// 通过 id 修改数据
func UpdateArticlesById(m *Articles) (err error) {
	o := orm.NewOrm()
	v := Articles{Id: m.Id}
	m.UpdatedAt = helper.NowFormat()
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m, "Title", "Category", "User", "Desc", "Content", "IndexImg", "IsRecommend", "IsTop", "Status", "Tags", "UpdatedAt", "Keywords"); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// 单个删除分类
func DeleteArticles(id int64) (err error) {
	o := orm.NewOrm()
	v := Articles{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Articles{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

//批量删除分类
func DeleteArticlesByIds(ids []string) (err error) {
	for _, val := range ids {
		id, err1 := strconv.ParseInt(val, 10, 64)
		if err1 != nil {
			err = err1
			return
		}
		err2 := DeleteArticles(id)
		if err2 != nil {
			err = err2
			continue
		}
	}
	return
}

//最新
func GetNewArticles(limit int) ([]Articles, error) {
	o := orm.NewOrm()
	var articles []Articles
	_, err := o.QueryTable(new(Articles)).OrderBy("id").Limit(limit).Filter("status", 1).RelatedSel().All(&articles, "Title", "Category", "User", "Desc", "IndexImg", "Tags", "CreatedAt")
	if err != nil {
		return articles, err
	}
	return articles, nil
}

//最热
func GetHotArticles(limit int) ([]Articles, error) {
	o := orm.NewOrm()
	var articles []Articles
	_, err := o.QueryTable(new(Articles)).OrderBy("ViewNum").Limit(limit).Filter("status", 1).RelatedSel().All(&articles, "Id", "Title", "Category", "User", "Desc", "IndexImg", "Tags", "CreatedAt")
	if err != nil {
		return articles, err
	}
	return articles, nil
}

//推荐
func GetTopArticles(limit int) ([]Articles, error) {
	o := orm.NewOrm()
	var articles []Articles
	_, err := o.QueryTable(new(Articles)).OrderBy("Id").Limit(limit).Filter("status", 1).Filter("IsTop", 1).RelatedSel().All(&articles, "Id", "Title", "Category", "User", "Desc", "IndexImg", "Tags", "CreatedAt")
	if err != nil {
		return articles, err
	}
	return articles, nil
}

//推荐
func GetTuiArticles(limit int) ([]Articles, error) {
	o := orm.NewOrm()
	var articles []Articles
	_, err := o.QueryTable(new(Articles)).OrderBy("Id").Limit(limit).Filter("status", 1).Filter("IsRecommend", 1).RelatedSel().All(&articles, "Id", "Title", "Category", "User", "Desc", "IndexImg", "Tags", "CreatedAt")
	if err != nil {
		return articles, err
	}
	return articles, nil
}

//点赞数
func UpdateZanNum(m *Articles) (err error) {
	o := orm.NewOrm()
	i := Articles{Id: m.Id}
	m.ZanNum += 1
	if err = o.Read(&i); err == nil {
		var num int64
		if num, err = o.Update(m, "ZanNum"); err == nil {
			fmt.Println(num)
		}
	}
	return
}

//阅读数
func UpdateViewNum(m *Articles) (err error) {
	o := orm.NewOrm()
	i := Articles{Id: m.Id}
	m.ViewNum += 1
	if err = o.Read(&i); err == nil {
		var num int64
		if num, err = o.Update(m, "ViewNum"); err == nil {
			fmt.Println(num)
		}
	}
	return
}
