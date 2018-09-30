package front

import (
	"blog/models"
	"html/template"
	"strconv"
)

type ArticleController struct {
	CommonController
}

func (c *ArticleController) Detail() {
	c.Layout = "layout/front/layout.html"
	c.Config()
	c.Tree()
	id := c.Ctx.Input.Param(":id")
	id64, _ := strconv.ParseInt(id, 0, 64)
	article, _ := models.GetArticlesById(id64)
	models.UpdateViewNum(article)
	c.Data["article_info"] = article
	c.Data["xsrf_token"] = template.HTML(c.XSRFToken())
	pre, _ := models.GetPreArticlesById(id64)
	if pre == nil {
		c.Data["pre"] = ""
	} else {
		c.Data["pre"] = "上一篇<br><a href='/article/" + strconv.FormatInt(pre.Id, 10) + "' rel='prev'>" + pre.Title + "</a>"
	}
	next, _ := models.GetNextArticlesById(id64)
	if next == nil {
		c.Data["next"] = ""
	} else {
		c.Data["next"] = "下一篇<br><a href='/article/" + strconv.FormatInt(pre.Id, 10) + "' rel='prev'>" + next.Title + "</a>"
	}

	c.LayoutSections = make(map[string]string)
	c.LayoutSections["HtmlHead"] = "layout/front/head.html"
	c.LayoutSections["HtmlMenu"] = "layout/front/menu.html"
	c.Data["body"] = "post-template-default single single-post postid-4495 single-format-standard nav_fixed list-comments-r p_indent comment-open site-layout-2"
	config, _ := models.GetSiteByType(1)
	c.Data["title"] = article.Title + " - " + config["app_name"]
	c.Data["keywords"] = config["home_keyword"]
	c.Data["description"] = article.Desc + " by " + config["app_name"]
	c.TplName = "front/article.tpl"
}

//点赞
func (c *ArticleController) Zan() {
	id := c.Ctx.Input.Param(":id")
	id64, _ := strconv.ParseInt(id, 0, 64)
	if _, err := models.GetZansByIpAndArticleId(c.Ctx.Request.RemoteAddr, id64); err == nil {
		c.Rsp(200, "您已经点过赞了!")
	} else {
		models.Zan(c.Ctx.Request.RemoteAddr, id64)
		article, _ := models.GetArticlesById(id64)
		models.UpdateZanNum(article)
		c.Rsp(200, "点赞成功!")
	}
}

//分类文章
func (c *ArticleController) ArticleCate() {
	c.Tree()
	c.Config()
	var fields []string
	var sortby []string
	var order []string
	var query = make(map[string]string)
	var limit int = 10
	var offset int
	order = append(order, "desc")
	sortby = append(sortby, "Id")
	pageno, _ := strconv.Atoi(c.Input().Get("page"))
	if pageno <= 0 {
		pageno = 1
	}
	offset = (pageno - 1) * limit
	cid := c.Input().Get("cid")
	if cid != "" {
		query["category_id__id"] = cid
		c.Data["cid"] = cid
	} else {
		c.Redirect("/", 301)
	}
	article_list, count, _ := models.GetAllArticles(query, fields, sortby, order, offset, limit)
	page := c.PageUtil(count, pageno, limit, article_list)
	c.Data["page"] = page
	c.Data["article_num"] = count
	cid64, _ := strconv.ParseInt(cid, 0, 64)
	category, _ := models.GetCategoriesById(cid64)
	c.Data["category"] = category
	c.Layout = "layout/front/layout.html"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["HtmlHead"] = "layout/front/head.html"
	c.LayoutSections["HtmlMenu"] = "layout/front/menu.html"
	config, _ := models.GetSiteByType(1)
	c.Data["body"] = "home blog nav_fixed list-comments-r site-layout-2"
	c.Data["app_name"] = config["app_name"]
	c.Data["title"] = category.Name + " - " + config["app_name"]
	c.Data["keywords"] = category.Keywords
	c.Data["description"] = category.Desc + " by " + config["app_name"]
	c.TplName = "front/article_cate.tpl"
}

//标签
func (c *ArticleController) ArticleTag() {
	c.Tree()
	c.Config()
	var fields []string
	var sortby []string
	var order []string
	var query = make(map[string]string)
	var limit int = 10
	var offset int
	order = append(order, "desc")
	sortby = append(sortby, "Id")
	pageno, _ := strconv.Atoi(c.Input().Get("page"))
	if pageno <= 0 {
		pageno = 1
	}
	offset = (pageno - 1) * limit
	tag := c.Input().Get("tag")
	color := c.Input().Get("color")
	if tag != "" {
		query["tags__contains"] = tag
		c.Data["tag"] = tag
		c.Data["color"] = color
	} else {
		c.Redirect("/", 301)
	}
	article_list, count, _ := models.GetAllArticles(query, fields, sortby, order, offset, limit)
	page := c.PageUtil(count, pageno, limit, article_list)
	c.Data["page"] = page
	c.Data["article_num"] = count
	c.Layout = "layout/front/layout.html"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["HtmlHead"] = "layout/front/head.html"
	c.LayoutSections["HtmlMenu"] = "layout/front/menu.html"
	config, _ := models.GetSiteByType(1)
	c.Data["body"] = "home blog nav_fixed list-comments-r site-layout-2"
	c.Data["app_name"] = config["app_name"]
	c.Data["title"] = tag + " - " + config["app_name"]
	c.Data["keywords"] = config["home_keyword"]
	c.Data["description"] = config["home_description"] + " by " + config["app_name"]
	c.TplName = "front/article_tag.tpl"
}
