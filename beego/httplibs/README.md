# httplib

## 导入包
```
    go get github.com/learninton/beegolibs/httplib
    
    import (
	    "github.com/learninton/beegolibs/httplib"
	)
```

## GET
```
    b,err := httplib.Get("req url")
    
    // 转换成自己的结构体
    var resData = 你的结构体
	json.Unmarshal(b, &resData)
```