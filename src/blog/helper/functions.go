package helper

import "time"

/*func AjaxJson(code int, msg string, count int, data interface{}) interface{} {
	return map[string]interface{}{"code": code, "msg": msg, "count": count, "list": data, "rel":true}
}*/

func NowFormat() (t string) {
	return time.Now().Format("2006-01-02 15:04:05")
}
