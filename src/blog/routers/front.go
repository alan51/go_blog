package routers

import (
	"blog/controllers/front"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &front.MainController{}, "get:Get")
	beego.Router("/article/:id", &front.ArticleController{}, "get:Detail")
	beego.Router("/article/zan/:id", &front.ArticleController{}, "post:Zan")
	beego.Router("/article_cate/", &front.ArticleController{}, "get:ArticleCate")
	beego.Router("/article_tag/", &front.ArticleController{}, "get:ArticleTag")
}
