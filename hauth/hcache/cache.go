package hcache

import (
	"crypto/sha1"
	"fmt"
	"github.com/astaxie/beego/cache"
	"time"
)

// beego cache modal
var bm, _ = cache.NewCache("memory", `{"interval":3600}`)

// get cached value by key.
func Get(key string) interface{} {
	return bm.Get(key)
}

// GetMulti is a batch version of GetDetails.
func GetMulti(keys []string) []interface{} {
	return bm.GetMulti(keys)
}

// set cached value with key and expire time.
func Put(key string, val interface{}, timeout time.Duration) error {
	return bm.Put(key, val, timeout)
}

// delete cached value by key.
func Delete(key string) error {
	return bm.Delete(key)
}

// increase cached int value by key, as a counter.
func Incr(key string) error {
	return bm.Incr(key)
}

// decrease cached int value by key, as a counter.
func Decr(key string) error {
	return bm.Decr(key)
}

// check if cached value exists or not.
func IsExist(key string) bool {
	return bm.IsExist(key)
}

// clear all cache.
func ClearAll() error {
	return bm.ClearAll()
}

// start gc routine based on config string settings.
func StartAndGC(config string) error {
	return bm.StartAndGC(config)
}

func GenKey(gpname string, keys ...string) string {
	sh := sha1.New()
	sh.Write([]byte(gpname))
	sh.Write([]byte("_join_"))
	for _, val := range keys {
		sh.Write([]byte(val))
		sh.Write([]byte("_join_"))
	}
	return fmt.Sprintf("%x", sh.Sum(nil))
}
