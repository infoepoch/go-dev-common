package restful

import (
	"encoding/json"

	"github.com/astaxie/beego"
)

type RestlibController struct {
	beego.Controller
}

/*
* 成功跳转
 */
func (this *RestlibController) Success(data interface{}) {
	this.Data["json"] = data
	this.ServeJSON()
	this.StopRun()
}

// 返回状态内容
func (this *RestlibController) SuccessList(list interface{}, sum_count int64) {
	data := make(map[string]interface{})
	data["list"] = list
	data["sum_count"] = sum_count
	this.Data["json"] = data
	this.ServeJSON()
	this.StopRun()
}

//失败返回
//{ "error_code": 错误码, "error_msg": "错误消息"}
func (this *RestlibController) Error(msg interface{}, code int) {
	data := make(map[string]interface{})
	data["errcode"] = code
	data["errmsg"] = msg
	this.Data["json"] = data
	this.ServeJSON()
	this.StopRun()
}

// 获取 post JSON 数据 转换成入参类型
func (this *RestlibController) GetPostJson(v interface{}) ([]byte, error) {
	b := this.Ctx.Input.RequestBody
	err := json.Unmarshal(b, &v)
	return b, err
}
