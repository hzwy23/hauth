package hcache

import (
	"time"

	"github.com/astaxie/beego/logs"
)

var relationCacheObj = NewRelationKey()

var expire = time.Minute * 720

// get cached value by key.
func Get(key string) interface{} {
	return relationCacheObj.cache.Get(key)
}

// GetMulti is a batch version of GetDetails.
func GetMulti(keys []string) []interface{} {
	return relationCacheObj.cache.GetMulti(keys)
}

func AddRelation(key string, value ...string) {
	relationCacheObj.SetRelatKey(key, value...)
}

func Set(key string, value interface{}) error {
	// 删除相关性缓存
	relKey, err := relationCacheObj.GetRelatKey(key)
	if err != nil {
		logs.Error(err)
		return err
	}
	for key, _ := range relKey {
		relationCacheObj.cache.Delete(key)
	}

	return relationCacheObj.cache.Put(key, value, expire)
}

// set cached value with key and expire time.
func Put(key string, val interface{}, timeout time.Duration) error {
	// 删除相关性缓存
	relKey, err := relationCacheObj.GetRelatKey(key)
	if err != nil {
		logs.Error(err)
		return err
	}
	for key, _ := range relKey {
		relationCacheObj.cache.Delete(key)
	}

	return relationCacheObj.cache.Put(key, val, timeout)
}

// delete cached value by key.
func Delete(key string) error {
	// 删除相关性缓存
	relKey, err := relationCacheObj.GetRelatKey(key)
	if err != nil {
		logs.Error(err)
		return err
	}
	for key, _ := range relKey {
		relationCacheObj.cache.Delete(key)
	}

	return relationCacheObj.cache.Delete(key)
}

// increase cached int value by key, as a counter.
func Incr(key string) error {
	// 删除相关性缓存
	relKey, err := relationCacheObj.GetRelatKey(key)
	if err != nil {
		logs.Error(err)
		return err
	}
	for key, _ := range relKey {
		relationCacheObj.cache.Delete(key)
	}

	return relationCacheObj.cache.Incr(key)
}

// decrease cached int value by key, as a counter.
func Decr(key string) error {
	// 删除相关性缓存
	relKey, err := relationCacheObj.GetRelatKey(key)
	if err != nil {
		logs.Error(err)
		return err
	}
	for key, _ := range relKey {
		relationCacheObj.cache.Delete(key)
	}

	return relationCacheObj.cache.Decr(key)
}

// check if cached value exists or not.
func IsExist(key string) bool {
	return relationCacheObj.cache.IsExist(key)
}

// clear all cache.
func ClearAll() error {
	return relationCacheObj.cache.ClearAll()
}

// start gc routine based on config string settings.
func StartAndGC(config string) error {
	return relationCacheObj.cache.StartAndGC(config)
}
