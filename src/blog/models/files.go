package models

import (
	"blog/helper"
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/astaxie/beego/orm"
	"os"
)

type Files struct {
	Id           int64
	Name         string `orm:"size(45)" valid:"Required;MaxSize:20;MinSize:2"`
	Domain       string `orm:"size(120)"`
	Url          string `orm:"size(150)"`
	CategoryName string `orm:"size(50)"`
	CreatedAt    string `orm:"null"`
}

func GetFiles() ([]orm.Params, error) {
	o := orm.NewOrm()
	file := new(Files)
	var files []orm.Params
	_, err := o.QueryTable(file).Values(&files)
	if err != nil {
		return files, err
	}
	return files, nil
}

//  新增文件
func AddFiles(m *Files) (id int64, err error) {
	o := orm.NewOrm()
	file := new(Files)
	file.Name = m.Name
	file.Domain = m.Domain
	file.Url = m.Url
	file.CategoryName = m.CategoryName
	file.CreatedAt = helper.NowFormat()
	id, err = o.Insert(file)
	return
}

// 详情
func GetFilesById(id int64) (v *Files, err error) {
	o := orm.NewOrm()
	v = &Files{Id: id}
	if err = o.QueryTable(new(Files)).Filter("Id", id).RelatedSel().One(v); err == nil {
		return v, nil
	}
	return nil, err
}

// 列表
func GetAllFiles(query map[string]string, fields []string, sortby []string, order []string,
	offset int, limit int) (ml []interface{}, count int, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Files))
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

	var l []Files
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

/*
func GetFileCategory()(cate []orm.Params, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Files))
	_, errs := qs.GroupBy("category_name").All(&cate)
	if errs != nil {
		err = errs
		return cate, err
	} else {
		return cate, nil
	}
}
*/

func GetFileCategory() ([]Files, error) {
	o := orm.NewOrm()
	file := new(Files)
	var files []Files
	_, err := o.QueryTable(file).GroupBy("category_name").All(&files, "CategoryName")
	if err != nil {
		return files, err
	}
	return files, nil
}

// 单个删除标签
func DeleteFiles(id int64) (err error) {
	o := orm.NewOrm()
	v := Files{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		file, errs := GetFilesById(id)
		if errs != nil {
			err = errs
			return
		}
		err := os.Remove("/www/wwwroot/blog/src/blog" + file.Url)
		if err == nil {
			if num, err = o.Delete(&Files{Id: id}); err == nil {
				fmt.Println("Number of records deleted in database:", num)
			}
		} else {
			return err
		}
	}
	return
}

//批量删除标签
func DeleteFilesByIds(ids []string) (err error) {
	for _, val := range ids {
		id, err1 := strconv.ParseInt(val, 10, 64)
		if err1 != nil {
			err = err1
			return
		}
		err2 := DeleteFiles(id)
		if err2 != nil {
			err = err2
			continue
		}
	}
	return
}
