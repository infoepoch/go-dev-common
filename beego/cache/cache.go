package cache

import (
	"bytes"
	"encoding/gob"
	"errors"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
	"github.com/astaxie/beego/logs"

	_ "github.com/astaxie/beego/cache/memcache"
	_ "github.com/astaxie/beego/cache/redis"
)

var (
	_cache cache.Cache
)

// init cached
func CacheInit() {

	cache_type := beego.AppConfig.String("cache_type")
	cache_config := beego.AppConfig.String("cache_config")

	logs.Info("cache type ", cache_type)

	var err error

	_cache, err = cache.NewCache(cache_type, cache_config)

	if err != nil {
		logs.Error("initialization cache", err)
	} else {
		logs.Info("initialization cache success")
	}
}

// get cached value by key.
func Get(key string, to interface{}) error {

	if _cache == nil {
		panic(errors.New("redis is nil"))
	}

	defer func() {
		if r := recover(); r != nil {
			//fmt.Println("get cache error caught: %v\n", r)
			logs.Error("get cache error caught", r)
			_cache = nil
		}
	}()

	data := _cache.Get(key)
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

// GetMulti is a batch version of Get.
func GetMulti(keys []string) []interface{} {
	return _cache.GetMulti(keys)
}

// set cached value with key and expire time.
func Put(key string, value interface{}, timeout int) error {
	data, err := GobEncode(value)
	if err != nil {
		logs.Error("set cache gob encode error", err)
		return err
	}
	if _cache == nil {
		return errors.New("redis is nil")
	}

	defer func() {
		if r := recover(); r != nil {
			logs.Error("set cache error caught", r)
			_cache = nil
		}
	}()

	timeouts := time.Duration(timeout) * time.Second
	err = _cache.Put(key, data, timeouts)
	if err != nil {
		logs.Error("set redis cache error，key:", key)
		return err
	} else {
		logs.Info("set redis cache success，key:", key)
		return nil
	}
}

// delete cached value by key.
func Delete(key string) error {
	if _cache == nil {
		return errors.New("cache is nil")
	}

	defer func() {
		if r := recover(); r != nil {
			logs.Error("delete cache error caught", r)
			_cache = nil
		}
	}()

	err := _cache.Delete(key)
	if err != nil {
		return errors.New("cache delete error")
	} else {
		logs.Info("cache delete success ", key)
		return nil
	}
}

// increase cached int value by key, as a counter.
func Incr(key string) error {
	return _cache.Incr(key)
}

// decrease cached int value by key, as a counter.
func Decr(key string) error {
	return _cache.Decr(key)
}

// check if cached value exists or not.
func IsExist(key string) bool {
	return _cache.IsExist(key)
}

// clear all cache.
func ClearAll() error {
	return _cache.ClearAll()
}

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
