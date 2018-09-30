package backend

import (
	"blog/models"
	"github.com/astaxie/beego"
)

type CommonController struct {
	beego.Controller
}

func init() {
	CheckAuth()
}

func (c *CommonController) Rsp(status int, msg string) {
	c.Data["json"] = c.AjaxJson(status, msg, 0, make(map[string]string))
	c.ServeJSON()
}

func (c *CommonController) AjaxJson(code int, msg string, count int, data interface{}) interface{} {
	return map[string]interface{}{"code": code, "msg": msg, "count": count, "list": data, "rel": true}
}

func (c *CommonController) GetTree(is_front int, pid int64) []models.MenuTree {
	nodes, _ := models.GetTree(is_front, pid)
	tree := make([]models.MenuTree, len(nodes))
	for k, v := range nodes {
		tree[k].Title = v.Name
		if k == 0 || k == 1 {
			tree[k].Spread = true
		} else {
			tree[k].Spread = false
		}
		tree[k].Href = v.Url
		tree[k].Icon = v.Icon
		tree[k].Id = v.Id
		children, _ := models.GetTree(is_front, v.Id)
		tree[k].Children = make([]models.MenuTree, len(children))
		for k1, v1 := range children {
			tree[k].Children[k1].Id = v1.Id
			tree[k].Children[k1].Title = v1.Name
			tree[k].Children[k1].Icon = v1.Icon
			tree[k].Children[k1].Href = v1.Url
		}
	}
	return tree
}

func (c *CommonController) GetTopTree(is_front int, pid int64) []models.MenuTreeTop {
	nodes, _ := models.GetTrees(pid)
	tree := make([]models.MenuTreeTop, len(nodes))
	for k, v := range nodes {
		tree[k].Title = v["Name"].(string)
		if k == 0 || k == 1 {
			tree[k].Spread = true
		} else {
			tree[k].Spread = false
		}
		tree[k].Href = v["Url"].(string)
		tree[k].Icon = v["Icon"].(string)
		tree[k].Id = v["Id"].(int64)
	}
	return tree
}

type Page struct {
	PageNo     int
	PageSize   int
	TotalPage  int
	TotalCount int
	FirstPage  bool
	LastPage   bool
	List       interface{}
}

func (c *CommonController) PageUtil(count int, pageNo int, pageSize int, list interface{}) Page {
	tp := count / pageSize
	if count%pageSize > 0 {
		tp = count/pageSize + 1
	}
	return Page{PageNo: pageNo, PageSize: pageSize, TotalPage: tp, TotalCount: count, FirstPage: pageNo == 1, LastPage: pageNo == tp, List: list}
}
