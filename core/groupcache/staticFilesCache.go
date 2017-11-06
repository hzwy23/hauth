package groupcache

import (
	"errors"
	"io/ioutil"
	"sync"

	"github.com/astaxie/beego/logs"
	"github.com/golang/groupcache"
)

// 这个包提供了缓存静态文件的方法
// 将系统中使用到的静态html文件缓存到内存中,
// 下次在打开这个页面时,就不需要从磁盘中读取,而是直接冲缓存中加载.
// example:
// groupcache.RegisterStaticFile(key,value)
// 每一个key只能注册一次,否则会panic.

var groupctx groupcache.Context
var lock = new(sync.RWMutex)

// 存储所有的静态文件
// key是索引
// vaule 是key的文件存储地址
var staticFile map[string]string = make(map[string]string)

// 注册静态文件信息
func RegisterStaticFile(key, value string) {
	lock.Lock()
	if _, ok := staticFile[key]; ok {
		lock.Unlock()
		panic(key + value + " 这个静态页面已经被注册了.")
	}
	staticFile[key] = value
	lock.Unlock()
}

func FileIsExist(key string) bool {
	lock.RLock()
	if _, ok := staticFile[key]; ok {
		lock.RUnlock()
		return true
	}
	lock.RUnlock()
	return false
}

func GetStaticFile(key string) ([]byte, error) {

	gp := groupcache.GetGroup("ASOFDATEHAUTH")
	if gp == nil {
		gp = groupcache.NewGroup("ASOFDATEHAUTH", 1<<28, groupcache.GetterFunc(func(ctx groupcache.Context, key string, dest groupcache.Sink) error {
			if filepath, ok := staticFile[key]; ok {
				logs.Debug("get html data from disk.")
				rst, _ := ioutil.ReadFile(filepath)
				return dest.SetBytes(rst)
			}
			return errors.New("filepath is not exists." + key)
		}))
	}

	var rst groupcache.ByteView
	err := gp.Get(groupctx, key, groupcache.ByteViewSink(&rst))
	if err != nil {
		goto DISK
	}
	logs.Debug("get data from cache.")
	return rst.ByteSlice(), err

DISK:
	if filepath, ok := staticFile[key]; ok {
		logs.Debug("get authority html data from disk.")
		rst, _ := ioutil.ReadFile(filepath)
		return rst, nil
	}
	return nil, errors.New("filepath is not exists." + key)
}
