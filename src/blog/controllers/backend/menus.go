package backend

import (
	"blog/models"
	"encoding/json"
	"errors"
	"github.com/astaxie/beego"
	"html/template"
	"strconv"
	"strings"
)

type MenusController struct {
	CommonController
}

func (c *MenusController) URLMapping() {
	c.Mapping("Index", c.Index)
	c.Mapping("DeleteSelect", c.DeleteSelect)
	c.Mapping("AddMenu", c.AddMenu)
	c.Mapping("EditMenu", c.EditMenu)

	c.Mapping("GetOne", c.GetOne)
	//c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)

}

//批量删除菜单
// @Title 批量删除菜单
// @Description 删除菜单
// @Param ids string true "删除的 id 集合"
// @Success 200 {object} 返回是否 cgong 成功
// @Failure 403
// @router /delete_select [post]
func (c *MenusController) DeleteSelect() {
	var ids []string
	if v := c.GetString("ids"); v != "" {
		ids = strings.Split(strings.TrimRight(v, ","), ",")
	}
	if len(ids) == 0 {
		c.Data["json"] = c.AjaxJson(403, "参数错误", 0, make(map[string]string))
	}
	err := models.DeleteMenuByIds(ids)
	if err != nil {
		c.Data["json"] = c.AjaxJson(201, err.Error(), 0, []interface{}{})
	} else {
		c.Data["json"] = c.AjaxJson(200, "删除成功", 0, make(map[string]string))
	}

	c.ServeJSON()
}

// 新增菜单页面
// @Title 菜单新增页面
// @Description 新增页面显示
// @Success 200 {object} 返回是否 cgong 成功
// @router /add_menu [get]
func (c *MenusController) AddMenu() {
	tree := c.GetTopTree(2, 0)
	c.Data["menus"] = tree
	category, _ := models.GetCategories()
	c.Data["category"] = category
	c.Data["xsrf_token"] = template.HTML(c.XSRFToken())
	c.TplName = beego.AppConfig.String("admin_view_path") + "menus/create.tpl"
}

// ajax 新增菜单接口
// @Title 新增菜单
// @Description xx
// @Success 200  {object} 返回是否成功
// @router /post_add_menu [post]
func (c *MenusController) PostAddMenu() {
	m := models.Menus{}
	if err := c.ParseForm(&m); err != nil {
		c.Rsp(403, err.Error())
		return
	}
	category_id, _ := c.GetInt64("CategoryId")
	if category_id > 0 {
		category := new(models.Categories)
		category.Id = category_id
		m.Category = category
	}
	_, err := models.AddMenus(&m)
	if err != nil {
		c.Rsp(403, err.Error())
		return
	} else {
		c.Rsp(200, "新增成功")
		return
	}
}

// 获取详情页面
// @title 获取详情页面
// @Description 详情页面
// @Failure 403 :id is empty
// @Success 200 {object} models.Menus
// @router /edit_menu/:id [get]
func (c *MenusController) EditMenu() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	v, err := models.GetMenusById(id)
	if v.Category == nil {
		c.Data["no_cate"] = true
	} else {
		c.Data["no_cate"] = false
	}
	if err != nil {
		return
	}
	c.Data["menu_info"] = v
	tree := c.GetTopTree(2, 0)
	c.Data["menus"] = tree
	category, _ := models.GetCategories()
	c.Data["category"] = category
	c.Data["xsrf_token"] = template.HTML(c.XSRFToken())
	c.TplName = beego.AppConfig.String("admin_view_path") + "menus/edit.tpl"
}

// 编辑菜单接口
// @Title 编辑菜单
// @Description 接口编辑菜单
// @Success 200 {object} 返回成功
// @router /post_edit_menu [post]
func (c *MenusController) PostEditMenu() {
	m := models.Menus{}
	if err := c.ParseForm(&m); err != nil {
		c.Rsp(403, err.Error())
		return
	}
	category_id, _ := c.GetInt64("CategoryId")
	if category_id > 0 {
		category := new(models.Categories)
		category.Id = category_id
		m.Category = category
	}
	err := models.UpdateMenusById(&m)
	if err != nil {
		c.Rsp(201, err.Error())
	} else {
		c.Rsp(200, "编辑成功")
	}
}

// GetOne ...
// @Title Get One
// @Description get Menus by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Menus
// @Failure 403 :id is empty
// @router /:id [get]
func (c *MenusController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	v, err := models.GetMenusById(id)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()
}

// 菜单首页
// @Title 菜单首页页面
// @Description 获取菜单分页数据
// @Param	query	query	string	false	"过滤. 例如. 过滤字段:过滤值,过滤字段:过滤值 ..."
// @Param	fields	query	string	false	"返回字段. 如. 字段一,字段二 ..."
// @Param	sortby	query	string	false	"根据什么字段排序. 如. 排序一,排序二 ..."
// @Param	order	query	string	false	"对应上面的,与每个排序字段对应的排序，如果是单个值，则适用于所有排序字段. 如. desc,asc ..."
// @Param	limit	query	string	false	"limit, 必须是整数"
// @Param	offset	query	string	false	"结果集的起始位置。 必须是整数"
// @Success 200 {object} 返回菜单列表或页面
// @Failure 403
// @router / [get]
func (c *MenusController) Index() {
	if c.IsAjax() == false {
		c.Data["xsrf_token"] = template.HTML(c.XSRFToken())
		//c.Layout = beego.AppConfig.String("admin_view_layout")
		c.TplName = beego.AppConfig.String("admin_view_path") + "menus/index.tpl"
	} else {
		var fields []string
		var sortby []string
		var order []string
		var query = make(map[string]string)
		var limit int = 10
		var offset int

		// fields: col1,col2,entity.col3
		if v := c.GetString("fields"); v != "" {
			fields = strings.Split(v, ",")
		}
		// limit: 10 (default is 10)
		if v, err := c.GetInt("pageSize"); err == nil {
			limit = v
		}
		// offset: 0 (default is 0)
		if v, err := c.GetInt("pageIndex"); err == nil {
			offset = (v - 1) * limit
		}
		// sortby: col1,col2
		if v := c.GetString("sort"); v != "" {
			sortby = strings.Split(v, ",")
		}
		// order: desc,asc
		if v := c.GetString("order"); v != "" {
			order = strings.Split(v, ",")
		}
		// query: k:v,k:v
		if v := c.GetString("query"); v != "" {
			for _, cond := range strings.Split(v, ",") {
				if len(cond) > 0 {
					kv := strings.SplitN(cond, ":", 2)
					if len(kv) != 2 {
						c.Data["json"] = errors.New("Error: invalid query key/value pair")
						c.ServeJSON()
						return
					}
					k, v := kv[0], kv[1]
					if len(v) > 0 {
						query[k] = v
					}
				}

			}
		}
		if len(order) <= 0 {
			order = append(order, "desc")
			sortby = append(sortby, "Id")
		}
		l, count, err := models.GetAllMenus(query, fields, sortby, order, offset, limit)
		if err != nil {
			c.Data["json"] = c.AjaxJson(1, err.Error(), count, make(map[string]string))
		} else {
			c.Data["json"] = c.AjaxJson(0, "成功!", count, l)
		}
		c.ServeJSON()
		return
	}
}

// Put ...
// @Title Put
// @Description update the Menus
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Menus	true		"body for Menus content"
// @Success 200 {object} models.Menus
// @Failure 403 :id is not int
// @router /:id [put]
func (c *MenusController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	v := models.Menus{Id: id}
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	if err := models.UpdateMenusById(&v); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// Delete ...
// @Title Delete
// @Description delete the Menus
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *MenusController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	if err := models.DeleteMenus(id); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}
