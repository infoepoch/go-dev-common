# uuid

## 引用方法：
```
go get github.com/infoepoch/go-dev-common/utils/uuid
```

## 使用方法
```golang
    s, e := uuid.GetUuid()
	if e != nil {
		t.Error(e.Error())
	}
	fmt.Println("uuid: ", s)
```