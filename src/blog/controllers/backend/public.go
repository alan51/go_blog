package backend

import (
	"blog/librarys"
	"blog/models"
	"errors"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"html/template"
)

type PublicController struct {
	beego.Controller
}

func (c *PublicController) AjaxJson(code int, msg string, count int, data interface{}) interface{} {
	return map[string]interface{}{"code": code, "msg": msg, "count": count, "list": data, "rel": true}
}

//登录
func (c *PublicController) Login() {
	if c.IsAjax() {
		username := c.GetString("username")
		password := c.GetString("password")
		if user, err := c.CheckLogin(username, password); err == nil {
			c.SetSession("userinfo", user)
			c.Data["json"] = c.AjaxJson(200, "登录成功", 0, make(map[string]string))
			c.ServeJSON()
			return
		} else {
			c.Data["json"] = c.AjaxJson(200, err.Error(), 0, make(map[string]string))
			c.ServeJSON()
			return
		}

	} else {
		userinfo := c.GetSession("userinfo")
		if userinfo != nil {
			c.Ctx.Redirect(301, "/admin/")
		}
		c.Data["xsrf_token"] = template.HTML(c.XSRFToken())
		c.TplName = beego.AppConfig.String("admin_view_path") + "main/login.tpl"
	}
}

//退出登录
func (c *PublicController) LogOut() {
	c.DelSession("userinfo")
	c.Ctx.Redirect(301, "/public/login")
}

func (c *PublicController) CheckLogin(username, password string) (user models.Users, err error) {
	user = models.GetUsersByName(username)
	if user.Id == 0 {
		return user, errors.New("用户不存在")
	}
	if user.Password != librarys.PwdHash(password) {
		return user, errors.New("账号或密码错误")
	}
	return user, nil
}

func CheckAuth() {

	var check = func(ctx *context.Context) {
		uinfo := ctx.Input.Session("userinfo")
		if uinfo == nil {
			ctx.Redirect(302, "/public/login")
			return
		} else {
			return
		}
	}
	beego.InsertFilter("/admin/*", beego.BeforeRouter, check)
}
