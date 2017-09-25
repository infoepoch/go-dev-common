---
name: cache 模块
sort: 2
---

# 缓存模块

beego 的 cache 模块是用来做数据缓存的，设计思路来自于 `database/sql`，目前支持 file、memcache、memory 和 redis 四种引擎，安装方式如下：

	go get github.com/infoepoch/go-dev-common/beego/cache


## 使用入门

首先在main.go 入口函数中 初始化，使得配置生效

    import (
        "github.com/infoepoch/go-dev-common/beego/cache"
    )

	func init() {
	    // cache_type 和 cache_config 参照引擎设置
    	cache.CacheInit(cache_type, cache_config)
    }

然后我们就可以使用bm增删改缓存：

    import (
        "github.com/infoepoch/go-dev-common/beego/cache"
    )

	cache.Put("astaxie", 1, 10*time.Second)
	cache.Get("astaxie")
	cache.IsExist("astaxie")
	cache.Delete("astaxie")

## 引擎设置

目前支持四种不同的引擎，接下来分别介绍这四种引擎如何设置：

- cache_type = memory

	配置信息如下所示，配置的信息表示 GC 的时间，表示每个 60s 会进行一次过期清理：

		cache_config = {"interval":60}
- cache_type = file

	配置信息如下所示，配置 `CachePath` 表示缓存的文件目录，`FileSuffix` 表示文件后缀，`DirectoryLevel` 表示目录层级，`EmbedExpiry` 表示过期设置

		cache_config = {"CachePath":"./cache","FileSuffix":".cache","DirectoryLevel":2,"EmbedExpiry":120}

- cache_type = redis

	配置信息如下所示，redis 采用了库 [redigo](https://github.com/garyburd/redigo/tree/master/redis):

		cache_config = {"key":"collectionName","conn":":6039","dbNum":"0","password":"thePassWord"}

	* key: Redis collection 的名称
	* conn: Redis 连接信息
	* dbNum: 连接 Redis 时的 DB 编号. 默认是0.
	* password: 用于连接有密码的 Redis 服务器.


- cache_type = memcache

	配置信息如下所示，memcache 采用了 [vitess的库](https://github.com/youtube/vitess/tree/master/go/memcache)，表示 memcache 的连接地址：

		cache_config = {"conn":"127.0.0.1:11211"}

## 开发自己的引擎

参考：

    github.com/astaxie/beego/cache/redis

beego cache 模块采用了接口的方式实现，因此用户可以很方便的实现接口，然后注册就可以实现自己的 Cache 引擎：

	type Cache interface {
		Get(key string) interface{}
        GetMulti(keys []string) []interface{}
		Put(key string, val interface{}, timeout time.Duration) error
		Delete(key string) error
		Incr(key string) error
		Decr(key string) error
		IsExist(key string) bool
		ClearAll() error
		StartAndGC(config string) error
	}

用户开发完毕在最后写类似这样的：

	func init() {
		Register("myowncache", NewOwnCache())
	}



## 我的使用案例

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

## cache val 建议加密

    /*
        Encode
        用gob进行数据编码
        todo 待迁移至公用库
    */
    func GobEncode(data interface{}) ([]byte, error) {
        buf := bytes.NewBuffer(nil)
        enc := gob.NewEncoder(buf)
        err := enc.Encode(data)
        if err != nil {
            return nil, err
        }
        return buf.Bytes(), nil
    }
    
    /*
        Decode
        用gob进行数据解码
        todo 待迁移至公用库
    */
    func GobDecode(data []byte, to interface{}) error {
        buf := bytes.NewBuffer(data)
        dec := gob.NewDecoder(buf)
        return dec.Decode(to)
    }