package caddy_plugins

import (
	"fmt"
	"sync"
	"time"
)

var (
	CacheObj = &Cache{
		make(map[string]*CacheData),
		&sync.Mutex{},
		NewMapWithExpire(nil),
	}
)

type CacheData struct {
	id   string
	data []byte
}

type Cache struct {
	dataMap  map[string]*CacheData
	mux      *sync.Mutex
	keyMap   *MapWithExpire
}

func (cache *Cache) SetData(key, id string, value []byte) error {

	val, ok := cache.keyMap.SetNx(key, id, 3 * time.Second)
	// 未获取到key的锁，或者没有id钥匙，直接返回
	if !ok || val.(string) != id {
		return nil
	}

	cache.mux.Lock()

	// id相同时append，不同或者不存在时创建
	if val, ok := cache.dataMap[key]; !ok || val.id != id {
		cache.dataMap[key] = &CacheData{
			id:   id,
			data: []byte{},
		}
	}
	cache.dataMap[key].data = append(cache.dataMap[key].data, value...)
	cache.mux.Unlock()

	return nil
}

func (c *Cache) GetKeyLen(key string) int {
	val, ok := c.GetData(key)
	if ok {
		return len(val)
	}
	return 0
}

func (cache *Cache) GetData(key string) ([]byte, bool) {

	val, ok := cache.dataMap[key]
	if ok {
		return val.data, ok
	} else {
		return nil, ok
	}
}



func (cache *Cache) GetLen() int {
	return len(cache.dataMap)
}

func (cache *Cache) GetKeys() string {

	var keys = ""
	for key, val := range cache.dataMap {
		keys += key + "," + fmt.Sprintf("(%d)", len(val.data))
	}
	return keys
}
