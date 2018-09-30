package controllers

import (
	"apiservice/models"
	"encoding/json"

	"github.com/astaxie/beego"
)

// 项目的增删改查接口文档
type ObjectController struct {
	beego.Controller
}

// @Title 创建
// @Description 创建一个项目
// @Param	body		body 	models.Object	true		"项目内容"
// @Success 200 {string} models.Object.Id
// @Failure 403 主体内容为空
// @router / [post]
func (o *ObjectController) Post() {
	var ob models.Object
	json.Unmarshal(o.Ctx.Input.RequestBody, &ob)
	objectid := models.AddOne(ob)
	o.Data["json"] = map[string]string{"ObjectId": objectid}
	o.ServeJSON()
}

// @Title 获取一条记录
// @Description 通过 id 获取项目的详情
// @Param	objectId		path 	string	true		"你想获取的项目详情的 id"
// @Success 200 {object} models.Object
// @Failure 403 :项目 id 不存在
// @router /:objectId [get]
func (o *ObjectController) Get() {
	objectId := o.Ctx.Input.Param(":objectId")
	if objectId != "" {
		ob, err := models.GetOne(objectId)
		if err != nil {
			o.Data["json"] = err.Error()
		} else {
			o.Data["json"] = ob
		}
	}
	o.ServeJSON()
}

// @Title 获取全部项目数据
// @Description 获取所以的项目
// @Success 200 {object} models.Object
// @Failure 403 :项目 id 不存在
// @router / [get]
func (o *ObjectController) GetAll() {
	obs := models.GetAll()
	o.Data["json"] = obs
	o.ServeJSON()
}

// @Title 更新项目操作
// @Description 更新项目
// @Param	objectId		path 	string	true		"项目 id"
// @Param	body		body 	models.Object	true		"内容主体"
// @Success 200 {object} models.Object
// @Failure 403 :项目 id 不存在
// @router /:objectId [put]
func (o *ObjectController) Put() {
	objectId := o.Ctx.Input.Param(":objectId")
	var ob models.Object
	json.Unmarshal(o.Ctx.Input.RequestBody, &ob)

	err := models.Update(objectId, ob.Score)
	if err != nil {
		o.Data["json"] = err.Error()
	} else {
		o.Data["json"] = "update success!"
	}
	o.ServeJSON()
}

// @Title 删除
// @Description 删除项目
// @Param	objectId		path 	string	true		"项目 id"
// @Success 200 {string} 删除成功
// @Failure 403 项目 id 为空
// @router /:objectId [delete]
func (o *ObjectController) Delete() {
	objectId := o.Ctx.Input.Param(":objectId")
	models.Delete(objectId)
	o.Data["json"] = "delete success!"
	o.ServeJSON()
}

