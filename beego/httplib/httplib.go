package httplib

import (
	"encoding/json"
	"github.com/astaxie/beego/httplib"
	"github.com/astaxie/beego/logs"
	"time"
)

var httpConnectTimeout time.Duration = 5
var httpReadWriteTimeout time.Duration = 5

func InitHttplib(_http_connect_timeout int, _http_read_write_timeout int) {
	httpConnectTimeout = time.Duration(_http_connect_timeout) * time.Second
	httpReadWriteTimeout = time.Duration(_http_read_write_timeout) * time.Second
}

func Get(reqUrl string) ([]byte, error) {
	logs.Info("request-get-url: " + reqUrl)

	b, err := httplib.Get(reqUrl).SetTimeout(httpConnectTimeout, httpReadWriteTimeout).Bytes()

	logs.Info("request-get-error: ", err)
	logs.Info("request-get-data: " + string(b))

	return b, err
}

// POST RAW JSON
func Post(reqUrl string, params interface{}) ([]byte, error) {
	logs.Info("request-post-url: " + reqUrl)
	log_params_str, _ := json.Marshal(params)
	logs.Info("request-post-params: ", string(log_params_str))

	req, _ := httplib.Post(reqUrl).SetTimeout(httpConnectTimeout, httpReadWriteTimeout).JSONBody(params)
	req.Header("Content-Type", "application/json;charset=utf-8")
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

// POST FROM
func PostForm(reqUrl string, params map[string]string) ([]byte, error) {
	logs.Info("request-post-from-url: " + reqUrl)


	req := httplib.Post(reqUrl)
	req.SetTimeout(httpConnectTimeout, httpReadWriteTimeout)
	for k, v := range params {
		req.Param(k, v)
	}
	b, err := req.Bytes()

	logs.Info("request-post-from-error: ", err)
	logs.Info("request-post-from-data: " + string(b))
	return b, err
}