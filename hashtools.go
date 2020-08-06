package caddy_plugins

import (
	"context"
	"sync"
	"time"
)

type ValueData struct {
	Data         interface{}
	AtExpireTime time.Time
}

type MapWithExpire struct {
	dataMap map[string]*ValueData
	mux     sync.Mutex
	ctx     context.Context
}

func NewMapWithExpire(ctx context.Context) *MapWithExpire {

	if ctx == nil {
		ctx = context.TODO()
	}

	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case <-time.After(1 * time.Second):
				// todo 这样写其实不好，参考定时任务的实现方式。参考时间轮实现方式： https://github.com/ouqiang/timewheel
				checkExpireKey()
			}
		}
	}()

	return &MapWithExpire{
		make(map[string]*ValueData),
		sync.Mutex{},
		ctx,
	}
}

// 模拟redis写法，内部的分布式锁
func (m *MapWithExpire) SetNx(key string, value interface{}, duration time.Duration) (interface{}, bool) {
	m.mux.Lock()
	defer m.mux.Unlock()

	now := time.Now()

	// 如果key存在并且未过期, set失败，返回当前值
	if val, ok := m.dataMap[key]; ok && now.Before(val.AtExpireTime){
		return val.Data, false
	}

	// 其他条件
	m.dataMap[key] = &ValueData{
		value,
		now.Add(duration),
	}

	return value, true
}

func checkExpireKey() {

}
