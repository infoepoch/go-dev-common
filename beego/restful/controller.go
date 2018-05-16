package restful

import (
	"encoding/json"

	"github.com/astaxie/beego"
)

type Controller struct {
	beego.Controller
}

type ResIo struct {
	Data   interface{} `json:"data"`
	Status int64       `json:"status"`
	Info   interface{} `json:"info"`
}

type ResListIo struct {
	Offset int64       `json:"offset"`
	Limit  int64       `json:"limit"`
	Total  int64       `json:"total"`
	Data   interface{} `json:"data"`
	Status int64       `json:"status"`
	Info   interface{} `json:"info"`
}

/*
* 成功跳转
 */
func (this *Controller) Success(data interface{}) {
	this.Data["json"] = ResIo{
		data,
		200,
		"Success",
	}
	this.ServeJSON()
	this.StopRun()
}

// 返回状态内容
func (this *Controller) SuccessList(io ResListIo) {
	this.Data["json"] = io
	this.ServeJSON()
	this.StopRun()
}

//失败返回
func (this *Controller) Error(r ResIo) {
	this.Data["json"] = r
	this.ServeJSON()
	this.StopRun()
}

// 获取 post JSON 数据 转换成入参类型
func (this *Controller) GetPostJson(v interface{}) ([]byte, error) {
	b := this.Ctx.Input.RequestBody
	err := json.Unmarshal(b, &v)
	return b, err
}
