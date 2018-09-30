package backend

import (
	"blog/models"
	"errors"
	"github.com/astaxie/beego"
	"html/template"
	"os"
	"strings"
	"time"
)

type FileUploadController struct {
	CommonController
}

func (f *FileUploadController) Prepare() {
	f.EnableXSRF = false
}

/*func (c *FileUploadController) URLMapping() {
	c.Mapping("Index", c.Index)
	c.Mapping("DeleteSelect", c.DeleteSelect)
	c.Mapping("UploadShop", c.UploadShop)
	c.Mapping("EditTag", c.EditTag)
}
*/
//批量删除图片
// @Title 批量删除图片
// @Description 删除图片
// @Param ids string true "删除的 id 集合"
// @Success 200 {object} 返回是否 cgong 成功
// @Failure 403
// @router /delete_select [post]
func (c *FileUploadController) DeleteSelect() {
	var ids []string
	if v := c.GetString("ids"); v != "" {
		ids = strings.Split(strings.TrimRight(v, ","), ",")
	}
	if len(ids) == 0 {
		c.Data["json"] = c.AjaxJson(403, "参数错误", 0, make(map[string]string))
	}
	err := models.DeleteFilesByIds(ids)
	if err != nil {
		c.Data["json"] = c.AjaxJson(201, err.Error(), 0, []interface{}{})
	} else {
		c.Data["json"] = c.AjaxJson(200, "删除成功", 0, make(map[string]string))
	}

	c.ServeJSON()
}

// ajax 新增图片接口
// @Title 新增图片
// @Description xx
// @Success 200  {object} 返回是否成功
// @router /upload [post]
func (c *FileUploadController) Upload() {
	path := "image/" + time.Now().Format("2006/01/02")
	upload_path := beego.AppConfig.String("uploadpath")
	full_path := upload_path + path
	err := os.MkdirAll(full_path, 0755)
	if err != nil {
		c.Data["json"] = c.AjaxJson(201, err.Error(), 0, []interface{}{})
		c.ServeJSON()
		return
	}
	f, h, _ := c.GetFile("image")
	host := c.Ctx.Request.Host
	err = c.SaveToFile("image", full_path+"/"+h.Filename)
	f.Close()
	if err != nil {
		c.Data["json"] = c.AjaxJson(201, err.Error(), 0, []interface{}{})
		c.ServeJSON()
		return
	}
	var image = new(models.Files)
	image.Name = h.Filename
	image.Domain = host
	catename := c.GetString("category_name")
	if len(catename) > 0 {
		image.CategoryName = catename
	} else {
		image.CategoryName = "未知分类"
	}
	image.Url = "/" + full_path + "/" + image.Name
	id, err := models.AddFiles(image)
	if err != nil {
		c.Data["json"] = c.AjaxJson(201, err.Error(), 0, []interface{}{})
		c.ServeJSON()
		return
	}
	image.Id = id
	c.Data["json"] = c.AjaxJson(200, "上传成功", 0, image)
	c.ServeJSON()
	return
}

// ajax editor.md 上传图片方法
// @Title 新增图片
// @Description xx
// @Success 200  {object} 返回是否成功
// @router /edit_upload [post]
func (c *FileUploadController) UploadEditor() {
	path := "image/" + time.Now().Format("2006/01/02")
	upload_path := beego.AppConfig.String("uploadpath")
	full_path := upload_path + path
	err := os.MkdirAll(full_path, 0755)
	if err != nil {
		c.Data["json"] = c.AjaxJson(201, err.Error(), 0, []interface{}{})
		c.ServeJSON()
		return
	}
	f, h, _ := c.GetFile("editormd-image-file")
	host := c.Ctx.Request.Host
	err = c.SaveToFile("editormd-image-file", full_path+"/"+h.Filename)
	f.Close()
	if err != nil {
		c.Data["json"] = c.AjaxJson(201, err.Error(), 0, []interface{}{})
		c.ServeJSON()
		return
	}
	var image = new(models.Files)
	image.Name = h.Filename
	image.Domain = host
	catename := c.GetString("category_name")
	if len(catename) > 0 {
		image.CategoryName = catename
	} else {
		image.CategoryName = "未知分类"
	}
	image.Url = "/" + full_path + "/" + image.Name
	id, err := models.AddFiles(image)
	if err != nil {
		c.Data["json"] = map[string]interface{}{"success": 0, "message": err.Error(), "url": ""}
		c.ServeJSON()
		return
	}
	image.Id = id
	c.Data["json"] = map[string]interface{}{"success": 1, "message": "上传成功", "url": image.Url}
	c.ServeJSON()
	return
}

// 图片库
// @Title 图片首页页面
// @Description 获取图片分页数据
// @Param	query	query	string	false	"过滤. 例如. 过滤字段:过滤值,过滤字段:过滤值 ..."
// @Param	fields	query	string	false	"返回字段. 如. 字段一,字段二 ..."
// @Param	sortby	query	string	false	"根据什么字段排序. 如. 排序一,排序二 ..."
// @Param	order	query	string	false	"对应上面的,与每个排序字段对应的排序，如果是单个值，则适用于所有排序字段. 如. desc,asc ..."
// @Param	limit	query	string	false	"limit, 必须是整数"
// @Param	offset	query	string	false	"结果集的起始位置。 必须是整数"
// @Success 200 {object} 返回图片列表或页面
// @Failure 403
// @router / [get]
func (c *FileUploadController) Index() {
	var fields []string
	var sortby []string
	var order []string
	var query = make(map[string]string)
	var limit int = 12
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
	sortby = append(sortby, "Id")
	// order: desc,asc
	if v := c.GetString("order"); v != "" {
		order = strings.Split(v, ",")
	}
	order = append(order, "desc")
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
	l, count, _ := models.GetAllFiles(query, fields, sortby, order, offset, limit)
	file_category, _ := models.GetFileCategory()
	c.Data["images"] = l
	c.Data["file_category"] = file_category
	c.Data["images_num"] = count
	pageIndex, _ := c.GetInt("pageIndex")
	c.Data["pageIndex"] = pageIndex
	cate_name := c.GetString("category_name")
	var current_url string
	if len(cate_name) > 0 {
		c.Data["category_name"] = cate_name
		current_url = c.Ctx.Request.Host + "/admin/file?category_name=" + cate_name + "&"
	} else {
		c.Data["category_name"] = ""
		current_url = c.Ctx.Request.Host + "/admin/file?"
	}
	if pageIndex > 0 {
		pre := pageIndex - 1
		if pre < 0 {
			pre = 1
		}
		c.Data["pre_url"] = current_url + "pageIndex=" + string(pre)
		c.Data["next_url"] = current_url + "pageIndex=" + string(pageIndex+1)
	} else {
		c.Data["pre_url"] = current_url + "pageIndex=1"
		c.Data["next_url"] = current_url + "pageIndex=2"
	}
	data := c.PageUtil(count, pageIndex, limit, l)
	c.Data["Page"] = data
	c.Data["xsrf_token"] = template.HTML(c.XSRFToken())
	//c.Layout = beego.AppConfig.String("admin_view_layout")
	c.TplName = beego.AppConfig.String("admin_view_path") + "files/index.tpl"
}

//  工具模板
// @Title 图片首页页面
// @Description 获取图片分页数据
// @Param	query	query	string	false	"过滤. 例如. 过滤字段:过滤值,过滤字段:过滤值 ..."
// @Param	fields	query	string	false	"返回字段. 如. 字段一,字段二 ..."
// @Param	sortby	query	string	false	"根据什么字段排序. 如. 排序一,排序二 ..."
// @Param	order	query	string	false	"对应上面的,与每个排序字段对应的排序，如果是单个值，则适用于所有排序字段. 如. desc,asc ..."
// @Param	limit	query	string	false	"limit, 必须是整数"
// @Param	offset	query	string	false	"结果集的起始位置。 必须是整数"
// @Success 200 {object} 返回图片列表或页面
// @Failure 403
// @router /upload [get]
func (c *FileUploadController) UploadShow() {
	var fields []string
	var sortby []string
	var order []string
	var query = make(map[string]string)
	var limit int = 12
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
	sortby = append(sortby, "Id")
	// order: desc,asc
	if v := c.GetString("order"); v != "" {
		order = strings.Split(v, ",")
	}
	order = append(order, "desc")
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
	l, count, _ := models.GetAllFiles(query, fields, sortby, order, offset, limit)
	file_category, _ := models.GetFileCategory()
	c.Data["images"] = l
	c.Data["file_category"] = file_category
	c.Data["images_num"] = count
	pageIndex, _ := c.GetInt("pageIndex")
	c.Data["pageIndex"] = pageIndex
	cate_name := c.GetString("category_name")
	var current_url string
	if len(cate_name) > 0 {
		c.Data["category_name"] = cate_name
		current_url = c.Ctx.Request.Host + "/admin/file?category_name=" + cate_name + "&"
	} else {
		c.Data["category_name"] = ""
		current_url = c.Ctx.Request.Host + "/admin/file?"
	}
	if pageIndex > 0 {
		pre := pageIndex - 1
		if pre < 0 {
			pre = 1
		}
		c.Data["pre_url"] = current_url + "pageIndex=" + string(pre)
		c.Data["next_url"] = current_url + "pageIndex=" + string(pageIndex+1)
	} else {
		c.Data["pre_url"] = current_url + "pageIndex=1"
		c.Data["next_url"] = current_url + "pageIndex=2"
	}
	data := c.PageUtil(count, pageIndex, limit, l)
	c.Data["Page"] = data
	c.Data["xsrf_token"] = template.HTML(c.XSRFToken())
	//c.Layout = beego.AppConfig.String("admin_view_layout")
	c.TplName = beego.AppConfig.String("admin_view_path") + "files/upload.tpl"
}
