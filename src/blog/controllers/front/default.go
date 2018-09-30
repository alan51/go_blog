package front

import (
	"blog/models"
	"github.com/astaxie/beego"
)

type MainController struct {
	CommonController
}

func (c *MainController) Get() {
	c.CommonData()
	c.Layout = "layout/front/layout.html"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["HtmlHead"] = "layout/front/head.html"
	c.LayoutSections["HtmlMenu"] = "layout/front/menu.html"
	config, _ := models.GetSiteByType(1)
	c.Data["body"] = "home blog nav_fixed list-comments-r site-layout-2"
	c.Data["app_name"] = config["app_name"]
	c.Data["title"] = config["app_name"] + " - " + config["title"]
	c.Data["keywords"] = config["home_keyword"]
	c.Data["description"] = config["home_description"] + " by " + config["app_name"]
	c.Data["qiniu_domain"] = beego.AppConfig.String("qiniu_domain")
	c.Data["qiniu_article"] = beego.AppConfig.String("qiniu_article")
	c.TplName = "front/index.tpl"
}
