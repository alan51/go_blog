package controllers

import (
	"apiservice/models"
	"encoding/json"

	"github.com/astaxie/beego"
)

// 用户的增删改查登录退出操作
type UserController struct {
	beego.Controller
}

// @Title 创建用户
// @Description 创建用户
// @Param	body		body 	models.User	true		"内容不能为空"
// @Success 200 {int} models.User.Id
// @Failure 403 创建参数为空
// @router / [post]
func (u *UserController) Post() {
	var user models.User
	json.Unmarshal(u.Ctx.Input.RequestBody, &user)
	uid := models.AddUser(user)
	u.Data["json"] = map[string]string{"uid": uid}
	u.ServeJSON()
}

// @Title 获取用户全部数据
// @Description 获取用户全部数据
// @Success 200 {object} models.User
// @router / [get]
func (u *UserController) GetAll() {
	users := models.GetAllUsers()
	u.Data["json"] = users
	u.ServeJSON()
}

// @Title 获取一条互数据
// @Description get user by uid
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.User
// @Failure 403 :用户 id 不能为空
// @router /:uid [get]
func (u *UserController) Get() {
	uid := u.GetString(":uid")
	if uid != "" {
		user, err := models.GetUser(uid)
		if err != nil {
			u.Data["json"] = err.Error()
		} else {
			u.Data["json"] = user
		}
	}
	u.ServeJSON()
}

// @Title 更新操作
// @Description 更新用户操作
// @Param	uid		path 	string	true		"The uid you want to update"
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {object} models.User
// @Failure 403 :uid 参数不是整数
// @router /:uid [put]
func (u *UserController) Put() {
	uid := u.GetString(":uid")
	if uid != "" {
		var user models.User
		json.Unmarshal(u.Ctx.Input.RequestBody, &user)
		uu, err := models.UpdateUser(uid, &user)
		if err != nil {
			u.Data["json"] = err.Error()
		} else {
			u.Data["json"] = uu
		}
	}
	u.ServeJSON()
}

// @Title 删除
// @Description 删除一个用户
// @Param	uid		path 	string	true		"填写你想删除的用户 uid"
// @Success 200 {string} 删除成功!
// @Failure 403 用户 uid 不能为空
// @router /:uid [delete]
func (u *UserController) Delete() {
	uid := u.GetString(":uid")
	models.DeleteUser(uid)
	u.Data["json"] = "delete success!"
	u.ServeJSON()
}

// @Title 登录
// @Description 登录系统用户
// @Param	username		query 	string	true		"登录用户的账号"
// @Param	password		query 	string	true		"登录用户的密码"
// @Success 200 {string} 登录成功
// @Failure 403 用户不存在
// @router /login [get]
func (u *UserController) Login() {
	username := u.GetString("username")
	password := u.GetString("password")
	if models.Login(username, password) {
		u.Data["json"] = "login success"
	} else {
		u.Data["json"] = "user not exist"
	}
	u.ServeJSON()
}

// @Title 退出登录
// @Description 清空登录用户的 session
// @Success 200 {string} 退出成功
// @router /logout [get]
func (u *UserController) Logout() {
	u.Data["json"] = "logout success"
	u.ServeJSON()
}

