package hcache

import (
	"errors"
	"sync"

	"github.com/astaxie/beego/cache"
)

type RelationKey struct {
	mp    map[string][]string
	lock  *sync.RWMutex
	cache cache.Cache
}

func NewRelationKey() *RelationKey {
	a := &RelationKey{
		mp:    make(map[string][]string),
		lock:  new(sync.RWMutex),
		cache: cache.NewMemoryCache(),
	}
	return a
}

// 设置某个key依赖的键
// 比如A依赖于B， 则 Set(A,B)
func (this *RelationKey) SetRelatKey(key string, value ...string) {
	this.lock.Lock()
	defer this.lock.Unlock()
	if _, ok := this.mp[key]; !ok {
		this.mp[key] = []string{}
	}

	if len(value) != 0 {
		for _, val := range value {
			if p, ok := this.mp[val]; ok {
				this.mp[val] = append(p, key)
			} else {
				this.mp[val] = []string{key}
			}
		}
	}
}

// 查询所有依赖于某个key的键集合
func (this *RelationKey) GetRelatKey(key string) (map[string]bool, error) {
	this.lock.Lock()
	defer this.lock.Unlock()
	var ret = make(map[string]bool)
	if val, ok := this.mp[key]; ok {
		this.recursive(val, ret)
		ret[key] = true
		return ret, nil
	} else {
		return ret, errors.New("key not found")
	}
}

// 获取所有的相关性key
func (this *RelationKey) recursive(data []string, ret map[string]bool) {
	for _, val := range data {
		if sublist, ok := this.mp[val]; ok {
			if _, yes := ret[val]; !yes {
				ret[val] = true
				this.recursive(sublist, ret)
			} else {
				continue
			}
		} else {
			ret[val] = true
		}
	}
}
