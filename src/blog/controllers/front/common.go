package front

import (
	"blog/models"
	"github.com/astaxie/beego"
	"strconv"
)

type CommonController struct {
	beego.Controller
}

func (c *CommonController) CommonData() {
	c.Article()
	c.Tree()
	c.Config()
}

func (c *CommonController) Article() {
	var fields []string
	var sortby []string
	var order []string
	order = append(order, "desc")
	sortby = append(sortby, "Id")
	var query = make(map[string]string)
	var limit int = 10
	var offset int
	pageno, _ := strconv.Atoi(c.Input().Get("page"))
	if pageno <= 0 {
		pageno = 1
	}
	offset = (pageno - 1) * limit
	if v := c.GetString("title"); v != "" {
		query["title__contains"] = v
		c.Data["search"] = v
	} else {
		c.Data["search"] = ""
	}
	article_list, count, _ := models.GetAllArticles(query, fields, sortby, order, offset, limit)
	page := c.PageUtil(count, pageno, limit, article_list)
	c.Data["page"] = page
	c.Data["article_num"] = count
}

func (c *CommonController) Tree() {
	tree := c.GetMenu(1, 0)
	c.Data["tree"] = tree
	current, _ := models.GetCategories()
	c.Data["current_tree"] = current
	//最新
	new_article, _ := models.GetNewArticles(5)
	c.Data["new_article"] = new_article
	c.Data["new_article_status"] = true
	if new_article == nil {
		c.Data["new_article_status"] = false
	}
	//最热
	hot_article, _ := models.GetHotArticles(5)
	c.Data["hot_article"] = hot_article
	c.Data["hot_article_status"] = true
	if hot_article == nil {
		c.Data["hot_article_status"] = false
	}
	//置顶
	top_article, _ := models.GetTopArticles(5)
	c.Data["top_article"] = top_article
	c.Data["top_article_status"] = true
	if top_article == nil {
		c.Data["top_article_status"] = false
	}
	//推荐
	tui_article, _ := models.GetTuiArticles(5)
	c.Data["tui_article"] = tui_article
	c.Data["tui_article_status"] = true
	if tui_article == nil {
		c.Data["tui_article_status"] = false
	}
	link, _ := models.GetLinks()
	c.Data["links"] = link
	c.Data["link_status"] = true
	if link == nil {
		c.Data["link_status"] = false
	}
}

func (c *CommonController) Config() {
	config, _ := models.GetSiteByType(1)
	c.Data["config"] = config
	c.Data["qiniu_domain"] = beego.AppConfig.String("qiniu_domain")
	c.Data["qiniu_article"] = beego.AppConfig.String("qiniu_article")
}

func (c *CommonController) Rsp(status int, msg string) {
	c.Data["json"] = c.AjaxJson(status, msg, 0, make(map[string]string))
	c.ServeJSON()
	c.StopRun()
}

func (c *CommonController) AjaxJson(code int, msg string, count int, data interface{}) interface{} {
	return map[string]interface{}{"code": code, "msg": msg, "count": count, "list": data, "rel": true}
}

func (c *CommonController) GetMenu(is_front int, pid int64) string {
	nodes, _ := models.GetTree(is_front, pid)
	var html string
	var categoryid string = c.Input().Get("cid")
	cid, _  := strconv.ParseInt(c.Input().Get("cid"), 10, 64)
	for k, v := range nodes {
		var id string = strconv.FormatInt(v.Id, 10)
		m_cate := v.Category
		var cate_id int64
		if m_cate != nil {
			cate_id = m_cate.Id
		}
		var name string = v.Name
		children, _ := models.GetTree(is_front, v.Id)
		var class string
		var drop string
		if len(children) > 0 {
			class += " menu-item-has-children menu-item-type-custom "
			drop += "<i class='fa fa-angle-down'></i>"
		} else {
			class += " menu-item-type-taxonomy "
		}
		if cid > 0 && v.Category != nil {
			categoryid_info,_ := models.GetCategoriesById(cid)
			if categoryid == strconv.FormatInt(cate_id, 10) || categoryid_info.ParentId == cate_id{
				class += " current-menu-item current_page_item "
			}
		} else if k == 0 && cid <=0 {
			class += " current-menu-item current_page_item "
		}
		html += "<li id='menu-item-" + id + "' class='menu-item menu-item-object-custom " + class + " menu-item-" + id + "'>"
		var icon string = v.Icon
		var ic string
		if icon != "" {
			ic += "<i class='" + icon + "'></i>"
		}
		var url string = v.Url
		if url == "" && cate_id > 0 {
			url = "/article_cate?cid=" + strconv.FormatInt(cate_id, 10)
		}
		html += "<a target='_black' href='" + url + "'>" + ic + name + drop + "</a>"

		if len(children) > 0 {
			html += "<ul class='sub-menu'>"
			for _, v1 := range children {
				var cid string = strconv.FormatInt(v1.Id, 10)
				var cname string = v1.Name
				cpid, _  := strconv.ParseInt(c.Input().Get("cid"), 10, 64)
				if cpid > 0 {
					categoryid_info,_ := models.GetCategoriesById(cpid)
					if v1.Category != nil && categoryid_info.Id == v1.Category.Id {
						html += "<li id='menu-item-" + cid + "' class='menu-item menu-item-type-custom menu-item-object-custom current-menu-item current_page_item menu-item-home menu-item-" + cid + "'>"
					} else {
						html += "<li id='menu-item-" + cid + "' class='menu-item menu-item-type-custom menu-item-object-custom menu-item-home menu-item-" + cid + "'>"
					}
				}


				var cicon string = v1.Icon
				var cic string
				if cicon != "" {
					cic += "<i class='" + cicon + "'></i>"
				}
				c_m_cate := v1.Category
				var c_cate_id int64
				if c_m_cate != nil {
					c_cate_id = c_m_cate.Id
				}
				var curl string = v1.Url
				if curl == "" && c_cate_id > 0 {
					curl = "/article_cate?cid=" + strconv.FormatInt(c_cate_id, 10)
				}
				html += "<a target='_black' href='" + curl + "'>" + cic + cname + "</a>"
			}
			html += "</ul>"
		}
		html += "</li>"

	}
	return html
}

func (c *CommonController) GetTopMenu(is_front int, pid int64) []models.MenuTreeTop {
	nodes, _ := models.GetTree(is_front, pid)
	tree := make([]models.MenuTreeTop, len(nodes))
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
	if pageNo <= 0 {
		pageNo = 1
	}
	if tp == 0 {
		tp = 1
	}
	return Page{PageNo: pageNo, PageSize: pageSize, TotalPage: tp, TotalCount: count, FirstPage: pageNo == 1, LastPage: pageNo == tp, List: list}
}
