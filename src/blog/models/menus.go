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

type Menus struct {
	Id        int64
	Name      string `orm:"size(45)" valid:"Required;MaxSize:20;MinSize:2"`
	Desc      string `orm:"size(128)"`
	Sort      int    `valid:"Required"`
	Pid       int64
	IsFront   int         `orm:"default(1)"`
	Url       string      `orm:"null"`
	Category  *Categories `orm:"rel(fk);null"`
	Icon      string      `orm:"null"`
	CreatedAt string      `orm:"null"`
	UpdatedAt string
}

type MenuTree struct {
	MenuTreeTop
	Children []MenuTree `json:"children"`
}
type MenuTreeTop struct {
	Id     int64  `json:"id"`
	Title  string `json:"title"`
	Icon   string `json:"icon"`
	Spread bool   `json:"spread"`
	Href   string `json:"href"`
}

func GetTree(is_front int, pid int64) ([]Menus, error) {
	o := orm.NewOrm()
	node := new(Menus)
	var nodes []Menus
	_, err := o.QueryTable(node).Filter("pid", pid).Filter("is_front", is_front).RelatedSel().OrderBy("Sort").All(&nodes)
	if err != nil {
		return nodes, err
	}
	return nodes, nil
}

func GetTrees(pid int64) ([]orm.Params, error) {
	o := orm.NewOrm()
	node := new(Menus)
	var nodes []orm.Params
	_, err := o.QueryTable(node).Filter("pid", pid).OrderBy("Sort").RelatedSel().Values(&nodes)
	if err != nil {
		return nodes, err
	}
	return nodes, nil
}

// AddMenus insert a new Menus into database and returns
// last inserted Id on success.
func AddMenus(m *Menus) (id int64, err error) {
	o := orm.NewOrm()
	Menu := new(Menus)
	Menu.Name = m.Name
	Menu.Desc = m.Desc
	Menu.Sort = int(m.Sort)
	Menu.Pid = int64(m.Pid)
	Menu.CreatedAt = helper.NowFormat()
	Menu.UpdatedAt = Menu.CreatedAt
	Menu.IsFront = m.IsFront
	Menu.Url = m.Url
	Menu.Icon = m.Icon
	Menu.Category = m.Category
	id, err = o.Insert(Menu)
	return
}

// GetMenusById retrieves Menus by Id. Returns error if
// Id doesn't exist
func GetMenusById(id int64) (v *Menus, err error) {
	o := orm.NewOrm()
	v = &Menus{Id: id}
	if err = o.QueryTable(new(Menus)).Filter("Id", id).RelatedSel().One(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllMenus retrieves all Menus matches certain condition. Returns empty list if
// no records exist
func GetAllMenus(query map[string]string, fields []string, sortby []string, order []string,
	offset int, limit int) (ml []interface{}, count int, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Menus))
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

	var l []Menus
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
					if fname == "CreatedAt" {
						m[fname] = ""
					} else {
						m[fname] = val.FieldByName(fname).Interface()
					}
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
func UpdateMenusById(m *Menus) (err error) {
	o := orm.NewOrm()
	v := Menus{Id: m.Id}
	m.UpdatedAt = helper.NowFormat()
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m, "Name", "Desc", "Pid", "UpdatedAt", "Sort", "IsFront", "Url", "Icon", "Category"); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteMenus deletes Menus by Id and returns error if
// the record to be deleted doesn't exist
func DeleteMenus(id int64) (err error) {
	o := orm.NewOrm()
	v := Menus{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Menus{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

//批量删除
func DeleteMenuByIds(ids []string) (err error) {
	for _, val := range ids {
		id, err1 := strconv.ParseInt(val, 10, 64)
		if err1 != nil {
			err = err1
			return
		}
		err2 := DeleteMenus(id)
		if err2 != nil {
			err = err2
			continue
		}
	}
	return
}
