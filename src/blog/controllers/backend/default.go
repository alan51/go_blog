package backend

import (
	"blog/models"
)

type MainController struct {
	CommonController
}

// @Title Post
// @Description create Categories
// @Param	body		body 	models.Categories	true		"body for Categories content"
// @Success 201 {int} models.Categories
// @Failure 403 body is empty
// @router /admin/main [get]
func (c *MainController) URLMapping() {
	c.Mapping("Get", c.Get)
}
func (c *MainController) Get() {
	c.Layout = "layout/backend/layout.html"
	c.Data["user"] = c.GetSession("userinfo")
	config, _ := models.GetSiteByType(2)
	c.Data["config"] = config
	c.TplName = "backend/main/main.tpl"
}

func (c *MainController) Index() {
	c.TplName = "backend/main/index.tpl"
}

func (c *MainController) Tree() {
	tree := c.GetTree(2, 0)
	c.Data["json"] = tree
	c.ServeJSON()
}
