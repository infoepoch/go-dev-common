package cache

import (
	"bytes"
	"encoding/gob"
	"errors"
	"time"

	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
	_ "github.com/astaxie/beego/cache/redis"
	"github.com/astaxie/beego/logs"
)

var cc cache.Cache

func InitRedis() {
	var err error

	defer func() {
		if r := recover(); r != nil {
			//fmt.Println("initial redis error caught: %v\n", r)
			cc = nil
		}
	}()

	redisHost := beego.AppConfig.String("redis_host")
	appName := beego.AppConfig.String("appname")
	logs.Info("redis host: " + redisHost)

	cc, err = cache.NewCache("redis", `{"key":"`+appName+`","conn":"`+redisHost+`"}`)

	if err != nil {
		logs.Error(err)
	} else {
		logs.Info("init redis success")
	}
}

func SetCache(key string, value interface{}, timeout int) error {
	data, err := GobEncode(value)
	if err != nil {
		return err
	}
	if cc == nil {
		return errors.New("redis is nil")
	}

	defer func() {
		if r := recover(); r != nil {
			//fmt.Println("set cache error caught: %v\n", r)
			cc = nil
		}
	}()

	timeouts := time.Duration(timeout) * time.Second
	err = cc.Put(key, data, timeouts)
	if err != nil {
		logs.Error("set redis cache error，key:", key)
		return err
	} else {
		logs.Info("set redis cache success，key:", key)
		return nil
	}
}

func GetCache(key string, to interface{}) error {

	if cc == nil {
		panic(errors.New("redis is nil"))
	}

	defer func() {
		if r := recover(); r != nil {
			//fmt.Println("get cache error caught: %v\n", r)
			cc = nil
		}
	}()

	data := cc.Get(key)
	if data == nil {
		return errors.New("redis cache not exist")
	}
	//fmt.Println(data)
	err := GobDecode(data.([]byte), to)
	if err != nil {
		logs.Error("get redis cache error", key, err)
	} else {
		logs.Info("get redis cache success", key)
	}

	return err
}

func DelCache(key string) error {
	if cc == nil {
		return errors.New("redis is nil")
	}

	defer func() {
		if r := recover(); r != nil {
			//fmt.Println("get cache error caught: %v\n", r)
			cc = nil
		}
	}()

	err := cc.Delete(key)
	if err != nil {
		return errors.New("redis cache delete error")
	} else {
		logs.Info("redis cache delete success ", key)
		return nil
	}
}

// --------------------
// Encode
// 用gob进行数据编码
// todo 待迁移至公用库
func GobEncode(data interface{}) ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	enc := gob.NewEncoder(buf)
	err := enc.Encode(data)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// -------------------
// Decode
// 用gob进行数据解码
// todo 待迁移至公用库
func GobDecode(data []byte, to interface{}) error {
	buf := bytes.NewBuffer(data)
	dec := gob.NewDecoder(buf)
	return dec.Decode(to)
}
