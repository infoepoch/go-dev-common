package utm

import (
	"net/url"

	"github.com/astaxie/beego"
	"github.com/infoepoch/go-dev-common/beego/restful"
)

// UtmIo UTM 结构体
type UtmIo struct {
	UtmSource   string `json:"utm_source"`   // 广告来源
	UtmMedium   string `json:"utm_medium"`   // 广告媒介
	UtmCampaign string `json:"utm_campaign"` // 广告名称
	UtmContent  string `json:"utm_content"`  // 广告内容
	UtmTerm     string `json:"utm_term"`     // 广告关键字
}

// GetUtm 获取
func GetUtm(c beego.Controller) {
	// 获取 url
	var utm_source := c.GetString("utm_source")
	var utm_medium := c.GetString("utm_medium")
	var utm_campaign := c.GetString("utm_campaign")
	var utm_content := c.GetString("utm_content")
	var utm_term := c.GetString("utm_term")

	// 设置到 session
	c.SetSession("utm_source", utm_source)
	c.SetSession("utm_medium", utm_medium)
	c.SetSession("utm_campaign", utm_campaign)
	c.SetSession("utm_content", utm_content)
	c.SetSession("utm_term", utm_term)
}

// GetUtmUrlStr 获取传递URL字符串
func GetUtmUrlStr(c restful.Controller) string {
	u := url.Values{}
	utm_source := c.GetSession("utm_source").(string)
	utm_medium := c.GetSession("utm_medium").(string)
	utm_campaign := c.GetSession("utm_campaign").(string)
	utm_content := c.GetSession("utm_content").(string)
	utm_term := c.GetSession("utm_term").(string)
	if utm_source != "" {
		u.Add("utm_source", utm_source)
	}
	if utm_medium != "" {
		u.Add("utm_medium", utm_medium)
	}
	if utm_campaign != "" {
		u.Add("utm_campaign", utm_campaign)
	}
	if utm_content != "" {
		u.Add("utm_content", utm_content)
	}
	if utm_term != "" {
		u.Add("utm_term", utm_term)
	}
	return u.Encode()
}
