package controllers

import (
	"github.com/astaxie/beego"
	"apiservice/models"
)
//文章接口文档
type ArticleController struct {
	beego.Controller
}
// @Title 获取用户全部数据
// @Description 获取用户全部数据
// @Success 200 {object} models.Article
// @router / [get]
func (c *ArticleController) getAll() {
	jsonData, _ := models.GetAllArticles()
	c.Data["json"] = jsonData
	c.ServeJSON()
}
