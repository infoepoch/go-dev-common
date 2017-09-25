---
name: httplib 模块
sort: 4
---

# 客户端请求

httplib 库主要用来模拟客户端发送 HTTP 请求，类似于 Curl 工具，支持 JQuery 类似的链式操作。使用起来相当的方便；通过如下方式进行安装：

	go get github.com/astaxie/beego/httplib
    go get github.com/infoepoch/go-dev-common/beego/httplib

首先在main.go 入口函数中 初始化，使得配置生效

    import (
	    "github.com/infoepoch/go-dev-common/beego/httplib"
	)

	func init() {
	    // 请求的超时时间和数据读取时间
    	httplib.InitHttplib(5, 5)
    }

然后我们就可以请求获取数据了

    b,err := httplib.Get("req url")
    
    // 转换成自己的结构体
    var resData = 你的结构体
	json.Unmarshal(b, &resData)

## 支持的方法对象

httplib 包里面支持如下的方法返回 []byte, error：

- Get(url string)
- Post(url string, parameter interface{})
- Put(url string, parameter interface{})