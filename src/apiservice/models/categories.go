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

type Categories struct {
	Id        int64
	Name      string `orm:"size(45)" valid:"Required;MaxSize:20;MinSize:2"`
	Icon      string `orm:"size(128)"`
	IsTop     int16  `valid:"Required"`
	Status    int16  `valid:"Required"`
	Url       string `orm:"size(120)"`
	Sort      int32  `valid:"Required"`
	ParentId  int64  `orm:"null"`
	CreatedAt string `orm:"null"`
	UpdatedAt string
	Articles  []*Articles `orm:"rel(m2m)"`
	Keywords  string      `orm:"size(255)"`
	Desc      string      `orm:"size(255)"`
}

func GetCategories() ([]Categories, error) {
	o := orm.NewOrm()
	cate := new(Categories)
	var cates []Categories
	_, err := o.QueryTable(cate).Filter("status", 1).Filter("is_top", 1).OrderBy("-Sort").RelatedSel().All(&cates)
	if err != nil {
		return cates, err
	}
	return cates, nil
}

//  新增网站分类
func AddCategories(m *Categories) (id int64, err error) {
	o := orm.NewOrm()
	category := new(Categories)
	category.Name = m.Name
	category.Icon = m.Icon
	category.Url = m.Url
	category.IsTop = int16(m.IsTop)
	category.Status = int16(m.Status)
	category.Sort = int32(m.Sort)
	category.CreatedAt = helper.NowFormat()
	category.UpdatedAt = category.CreatedAt
	category.Keywords = m.Keywords
	category.Desc = m.Desc
	category.ParentId = m.ParentId
	id, err = o.Insert(category)
	return
}

// 详情
func GetCategoriesById(id int64) (v *Categories, err error) {
	o := orm.NewOrm()
	v = &Categories{Id: id}
	if err = o.QueryTable(new(Categories)).Filter("Id", id).RelatedSel().One(v); err == nil {
		return v, nil
	}
	return nil, err
}

// 列表
func GetAllCategories(query map[string]string, fields []string, sortby []string, order []string,
	offset int, limit int) (ml []interface{}, count int, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Categories))
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

	var l []Categories
	qs = qs.OrderBy(sortFields...).RelatedSel()
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
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
func UpdateCategoriesById(m *Categories) (err error) {
	o := orm.NewOrm()
	v := Categories{Id: m.Id}
	m.UpdatedAt = helper.NowFormat()
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m, "Name", "Icon", "IsTop", "Sort", "ParentId", "Status", "UpdatedAt", "Keywords", "Desc"); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// 单个删除分类
func DeleteCategories(id int64) (err error) {
	o := orm.NewOrm()
	v := Categories{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Categories{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

//批量删除分类
func DeleteCategoriesByIds(ids []string) (err error) {
	for _, val := range ids {
		id, err1 := strconv.ParseInt(val, 10, 64)
		if err1 != nil {
			err = err1
			return
		}
		err2 := DeleteCategories(id)
		if err2 != nil {
			err = err2
			continue
		}
	}
	return
}
