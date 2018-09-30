package main

import (
	_ "blog/routers"

	"blog/models"
	"encoding/gob"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"strings"
)

func main() {
	logs.SetLogger("console")
	logs.SetLogger(logs.AdapterFile, `{"filename":"logs/project.log","level":7,"maxlines":0,"maxsize":0,"daily":true,"maxdays":10}`)
	beego.SetStaticPath("/swagger", "swagger")
	beego.AddFuncMap("StrToSplit", StrToSplit)
	beego.Run()
}

func init() {
	beego.BConfig.WebConfig.Session.SessionProvider = "file"
	beego.BConfig.WebConfig.Session.SessionProviderConfig = "./tmp"
	gob.Register(models.Users{})
}

func StrToSplit(in string) (out string) {
	tags := strings.Split(in, "\n")
	for k, v:=range tags {
		var color string
		switch k {
			case 0:
				color = "727577"
			case 1:
				color = "a77676"
			case 2:
				color = "5d5252"
			case 3:
				color = "5a5a7d"
			case 4:
				color = "bf8835"
			case 5:
				color = "929266"
			case 6:
				color ="91849e"
			default:
				color = "727577"
		}
		out += "<a class='tag' target='_black' style='background:#"+color+"' title='"+v+"' href='/article_tag?tag="+v+"&color="+color+"' rel='tag'>"+v+"</a>"
	}
	return
}