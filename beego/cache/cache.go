package cache

import (
	"time"

	"github.com/astaxie/beego/cache"
	"github.com/astaxie/beego/logs"

	"errors"

	_ "github.com/astaxie/beego/cache/memcache"
	_ "github.com/astaxie/beego/cache/redis"
)

var (
	_cache cache.Cache
)

// init cached
func InitCache(_type string, _config string) error {

	logs.Info("cache type ", _type)

	var err error

	_cache, err = cache.NewCache(_type, _config)

	if err != nil {
		logs.Error("initialization cache", err)
	} else {
		logs.Info("initialization cache success")
	}

	return err
}

// Get cached value by key.
func Get(key string) interface{} {
	if _cache == nil {
		logs.Error("get cache is nil")
		return errors.New("cache is nil")
	}

	defer func() {
		if r := recover(); r != nil {
			logs.Error("get cache error caught", r)
			_cache = nil
		}
	}()

	val := _cache.Get(key)
	return val
}

// GetMulti is a batch version of Get.
func GetMulti(keys []string) []interface{} {
	if _cache == nil {
		logs.Error("get multi cache is nil")
		return nil
	}

	defer func() {
		if r := recover(); r != nil {
			logs.Error("get multi cache error caught", r)
			_cache = nil
		}
	}()

	return _cache.GetMulti(keys)
}

// Put set cached value with key and expire time.
func Put(key string, value interface{}, timeout int64) error {
	if _cache == nil {
		return errors.New("put cache is nil")
	}

	defer func() {
		if r := recover(); r != nil {
			logs.Error("put cache error caught", r)
			_cache = nil
		}
	}()

	tot := time.Duration(timeout) * time.Second

	return _cache.Put(key, value, tot)
}

// Delete cached value by key.
func Delete(key string) error {
	if _cache == nil {
		return errors.New("delete cache is nil")
	}

	defer func() {
		if r := recover(); r != nil {
			logs.Error("delete cache error caught", r)
			_cache = nil
		}
	}()

	err := _cache.Delete(key)
	if err != nil {
		return errors.New("delete cache error")
	} else {
		logs.Info("delete cache success ", key)
		return nil
	}
}

// Incr increase cached int value by key, as a counter.
func Incr(key string) error {
	if _cache == nil {
		logs.Error("incr cache is nil")
		return errors.New("cache is nil")
	}

	defer func() {
		if r := recover(); r != nil {
			logs.Error("incr cache error caught", r)
			_cache = nil
		}
	}()

	return _cache.Incr(key)
}

// Decr decrease cached int value by key, as a counter.
func Decr(key string) error {
	if _cache == nil {
		logs.Error("decr cache is nil")
		return errors.New("cache is nil")
	}

	defer func() {
		if r := recover(); r != nil {
			logs.Error("decr cache error caught", r)
			_cache = nil
		}
	}()

	return _cache.Decr(key)
}

// IsExist check if cached value exists or not.
func IsExist(key string) bool {
	if _cache == nil {
		logs.Error("is_exist cache is nil")
		return false
	}

	defer func() {
		if r := recover(); r != nil {
			logs.Error("is_exist cache error caught", r)
			_cache = nil
		}
	}()

	return _cache.IsExist(key)
}

// ClearAll cache.
func ClearAll() error {
	if _cache == nil {
		logs.Error("clear_all cache is nil")
		return errors.New("clear_all cache is nil")
	}

	defer func() {
		if r := recover(); r != nil {
			logs.Error("clear_all cache error caught", r)
			_cache = nil
		}
	}()

	return _cache.ClearAll()
}
