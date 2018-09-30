package models

import (
	"blog/helper"
	"errors"
	"fmt"
	"github.com/astaxie/beego/orm"
	"reflect"
	"strconv"
	"strings"
)

type Tags struct {
	Id        int64
	Name      string `orm:"size(45)" valid:"Required;MaxSize:20;MinSize:2"`
	CreatedAt string `orm:"null"`
	UpdatedAt string
}

func GetTags() ([]orm.Params, error) {
	o := orm.NewOrm()
	cate := new(Tags)
	var cates []orm.Params
	_, err := o.QueryTable(cate).Values(&cates)
	if err != nil {
		return cates, err
	}
	return cates, nil
}

//  新增网站标签
func AddTags(m *Tags) (id int64, err error) {
	o := orm.NewOrm()
	category := new(Tags)
	category.Name = m.Name
	category.CreatedAt = helper.NowFormat()
	category.UpdatedAt = category.CreatedAt
	id, err = o.Insert(category)
	return
}

// 详情
func GetTagsById(id int64) (v *Tags, err error) {
	o := orm.NewOrm()
	v = &Tags{Id: id}
	if err = o.QueryTable(new(Tags)).Filter("Id", id).RelatedSel().One(v); err == nil {
		return v, nil
	}
	return nil, err
}

// 列表
func GetAllTags(query map[string]string, fields []string, sortby []string, order []string,
	offset int, limit int) (ml []interface{}, count int, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Tags))
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

	var l []Tags
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
func UpdateTagsById(m *Tags) (err error) {
	o := orm.NewOrm()
	v := Tags{Id: m.Id}
	m.UpdatedAt = helper.NowFormat()
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m, "Name", "UpdatedAt"); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// 单个删除标签
func DeleteTags(id int64) (err error) {
	o := orm.NewOrm()
	v := Tags{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Tags{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

//批量删除标签
func DeleteTagsByIds(ids []string) (err error) {
	for _, val := range ids {
		id, err1 := strconv.ParseInt(val, 10, 64)
		if err1 != nil {
			err = err1
			return
		}
		err2 := DeleteTags(id)
		if err2 != nil {
			err = err2
			continue
		}
	}
	return
}
