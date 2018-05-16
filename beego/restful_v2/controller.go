package restful_v2

import (
	"encoding/json"

	"github.com/astaxie/beego"
)

type Controller struct {
	beego.Controller
}

/*
* 成功跳转
 */
func (this *Controller) Success(data interface{}) {
	this.Data["json"] = data
	this.ServeJSON()
	this.StopRun()
}

// 返回状态内容
func (this *Controller) SuccessList(list interface{}, total int64, offset int64, limit int64) {
	data := make(map[string]interface{})
	data["offset"] = offset
	data["limit"] = limit
	data["list"] = list
	data["total"] = total
	this.Data["json"] = data
	this.ServeJSON()
	this.StopRun()
}

//失败返回
//{ "error_code": 错误码, "error_msg": "错误消息"}
func (this *Controller) Error(message interface{}, code int) {
	data := make(map[string]interface{})
	data["errorCode"] = code
	data["message"] = message
	this.Data["json"] = data
	this.ServeJSON()
	this.StopRun()
}

// 获取 post JSON 数据 转换成入参类型
func (this *Controller) GetPostJson(v interface{}) ([]byte, error) {
	b := this.Ctx.Input.RequestBody
	err := json.Unmarshal(b, &v)
	return b, err
}
