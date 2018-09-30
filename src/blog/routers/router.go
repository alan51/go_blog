package routers

import (
	"blog/controllers/backend"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/public/login", &backend.PublicController{}, "*:Login")
	beego.Router("/public/logout", &backend.PublicController{}, "*:LogOut")
	ns := beego.NewNamespace("/admin",
		beego.NSRouter("/", &backend.MainController{}),
		beego.NSRouter("/index", &backend.MainController{}, "get:Index"),
		beego.NSRouter("/tree", &backend.MainController{}, "get:Tree"),
		beego.NSNamespace("/menu",
			beego.NSInclude(&backend.MenusController{}),
		),
		beego.NSNamespace("/site",
			beego.NSInclude(&backend.SiteController{}),
		),
		beego.NSNamespace("/category",
			beego.NSInclude(&backend.CategoryController{}),
		),
		beego.NSNamespace("/tag",
			beego.NSInclude(&backend.TagController{}),
		),
		beego.NSNamespace("/file",
			beego.NSInclude(&backend.FileUploadController{}),
		),
		beego.NSNamespace("/article",
			beego.NSInclude(&backend.ArticleController{}),
		),
		beego.NSNamespace("/link",
			beego.NSInclude(&backend.LinkController{}),
		),
		beego.NSNamespace("/user",
			beego.NSInclude(&backend.UserController{}),
		),
	)
	beego.AddNamespace(ns)
}
