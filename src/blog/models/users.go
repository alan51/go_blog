package models

import (
	"blog/helper"
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/astaxie/beego/orm"

	"blog/librarys"
)

type Users struct {
	Id        int64
	UserName  string `orm:"size(45)" valid:"Required;MaxSize:20;MinSize:2;unique"`
	NickName  string `orm:"size(45)" valid:"Required;MaxSize:20;MinSize:2"`
	Password  string `orm:"size(125)"`
	Img       string `orm:"size(128)"`
	Status    int16  `valid:"Required"`
	Url       string `orm:"size(120)"`
	CreatedAt string `orm:"null"`
	UpdatedAt string
	Articles  []*Articles `orm:"rel(m2m)"`
}

func GetUsers() ([]orm.Params, error) {
	o := orm.NewOrm()
	cate := new(Users)
	var cates []orm.Params
	_, err := o.QueryTable(cate).Filter("status", 1).OrderBy("-Sort").Values(&cates)
	if err != nil {
		return cates, err
	}
	return cates, nil
}

//  新增友情链接
func AddUsers(m *Users) (id int64, err error) {
	o := orm.NewOrm()
	link := new(Users)
	link.UserName = m.UserName
	link.NickName = m.NickName
	link.Password = librarys.StrToMd5(m.Password)
	link.Img = m.Img
	link.Url = m.Url
	link.Status = int16(m.Status)
	link.CreatedAt = helper.NowFormat()
	link.UpdatedAt = link.CreatedAt
	id, err = o.Insert(link)
	return
}

// 详情
func GetUsersById(id int64) (v *Users, err error) {
	o := orm.NewOrm()
	v = &Users{Id: id}
	if err = o.QueryTable(new(Users)).Filter("Id", id).RelatedSel().One(v); err == nil {
		return v, nil
	}
	return nil, err
}

//根据名字获取用户
func GetUsersByName(username string) (user Users) {
	o := orm.NewOrm()
	user = Users{UserName: username}
	o.Read(&user, "UserName")
	return user
}

// 列表
func GetAllUsers(query map[string]string, fields []string, sortby []string, order []string,
	offset int, limit int) (ml []interface{}, count int, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Users))
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

	var l []Users
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
func UpdateUsersById(m *Users) (err error) {
	o := orm.NewOrm()
	v := Users{Id: m.Id}
	m.UpdatedAt = helper.NowFormat()
	if err = o.Read(&v); err == nil {
		var num int64
		m.Password = librarys.StrToMd5(m.Password)
		if num, err = o.Update(m, "UserName", "NickName", "Password", "Img", "Url", "Status", "UpdatedAt"); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// 单个删除友情链接
func DeleteUsers(id int64) (err error) {
	o := orm.NewOrm()
	v := Users{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Users{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

//批量删除友情链接
func DeleteUsersByIds(ids []string) (err error) {
	for _, val := range ids {
		id, err1 := strconv.ParseInt(val, 10, 64)
		if err1 != nil {
			err = err1
			return
		}
		err2 := DeleteUsers(id)
		if err2 != nil {
			err = err2
			continue
		}
	}
	return
}
