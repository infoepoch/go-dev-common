# Redis cache 使用方法
----------
## cache 用法
```golang
// 引用
import "github.com/infoepoch/go-dev-common/beego/utils/cache"

// 调用方法初始化
cache.InitRedis()

// 调用方法

// 获取 JsapiTicket
func GetJsapiTicket() (server.JsApiToken, error) {
	// 解决并发多同时访问
	var mutex sync.Mutex
	var cacheName = setting.AppName + "-JsApiToken-" + conf.CORP_ID
	mutex.Lock()
	var token server.JsApiToken
    // 获取缓存
	err := cache.GetCache(cacheName, &token) 
	if err != nil {
		accessToken, _ := GetAccessToken()
		token, err = server.GetJsApiToken(accessToken.AccessToken)
		// 设置缓存
		cache.SetCache(cacheName, token, 60)
	}
	//判断缓存读取缓存
	mutex.Unlock()
	return token, err
}

```