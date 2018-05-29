package restful

import (
	"encoding/json"

	"github.com/astaxie/beego"
	"net/url"
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

type UtmIo struct {
	UtmSource   string `json:"utm_source"`   // 广告来源
	UtmMedium   string `json:"utm_medium"`   // 广告媒介
	UtmCampaign string `json:"utm_campaign"` // 广告名称
	UtmContent  string `json:"utm_content"`  // 广告内容
	UtmTerm     string `json:"utm_term"`     // 广告关键字
}

func (c *Controller) GetUtm() {
	// 获取 url
	utm_source := c.GetString("utm_source")
	utm_medium := c.GetString("utm_medium")
	utm_campaign := c.GetString("utm_campaign")
	utm_content := c.GetString("utm_content")
	utm_term := c.GetString("utm_term")

	// 设置到 session
	c.SetSession("utm_source", utm_source)
	c.SetSession("utm_medium", utm_medium)
	c.SetSession("utm_campaign", utm_campaign)
	c.SetSession("utm_content", utm_content)
	c.SetSession("utm_term", utm_term)
}

// 获取传递URL字符串
func (c *Controller) GetUtmUrlStr() string {
	u := url.Values{}
	utm_source := c.GetSession("utm_source")
	if utm_source != nil && utm_source.(string) != "" {
		u.Add("utm_source", utm_source.(string))
	}

	utm_medium := c.GetSession("utm_medium")
	if utm_medium != nil && utm_medium.(string) != "" {
		u.Add("utm_medium", utm_medium.(string))
	}

	utm_campaign := c.GetSession("utm_campaign")
	if utm_campaign != nil && utm_campaign.(string) != "" {
		u.Add("utm_campaign", utm_campaign.(string))
	}

	utm_content := c.GetSession("utm_content")
	if utm_content != nil && utm_content != "" {
		u.Add("utm_content", utm_content.(string))
	}

	utm_term := c.GetSession("utm_term")
	if utm_term != nil && utm_term.(string) != "" {
		u.Add("utm_term", utm_term.(string))
	}
	return u.Encode()
}
