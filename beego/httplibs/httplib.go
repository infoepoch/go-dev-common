package httplib

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
	"github.com/astaxie/beego/logs"
	"time"
)

var httpConnectTimeout time.Duration
var httpReadWriteTimeout time.Duration

func init() {
	hct, _ := beego.AppConfig.Int64("httpconnecttimeout")
	hrwt, _ := beego.AppConfig.Int64("httpreadwritetimeout")

	httpConnectTimeout = time.Duration(hct) * time.Second
	httpReadWriteTimeout = time.Duration(hrwt) * time.Second
	// TODO 运行 go test  取消 21行和22行的注释
	//httpConnectTimeout = time.Duration(5) * time.Second
	//httpReadWriteTimeout = time.Duration(5) * time.Second

}

func Get(reqUrl string) ([]byte, error) {
	logs.Info("request-get-url: " + reqUrl)

	b, err := httplib.Get(reqUrl).SetTimeout(httpConnectTimeout, httpReadWriteTimeout).Bytes()

	logs.Info("request-get-error: ", err)
	logs.Info("request-get-data: " + string(b))

	return b, err
}

func Post(reqUrl string, params interface{}) ([]byte, error) {
	logs.Info("request-post-url: " + reqUrl)
	log_params_str, _ := json.Marshal(params)
	logs.Info("request-post-params: ", string(log_params_str))

	req, _ := httplib.Post(reqUrl).SetTimeout(httpConnectTimeout, httpReadWriteTimeout).JSONBody(params)
	b, err := req.Bytes()

	logs.Info("request-post-error: ", err)
	logs.Info("request-post-data: " + string(b))
	return b, err
}

func Put(reqUrl string, params interface{}) ([]byte, error) {
	logs.Info("request-put-url: " + reqUrl)
	log_params_str, _ := json.Marshal(params)
	logs.Info("request-put-params: " + string(log_params_str))

	req, _ := httplib.Put(reqUrl).SetTimeout(httpConnectTimeout, httpReadWriteTimeout).JSONBody(params)
	b, err := req.Bytes()

	logs.Info("request-put-error: ", err)
	logs.Info("request-put-data: " + string(b))
	return b, err
}
