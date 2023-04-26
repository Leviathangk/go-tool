// Package gsync 并发 map，但是获取数据后需要转换，或者使用 sync.Map
package gsync

import "sync"

// SyncMap 并发 map
type SyncMap struct {
	Map  map[interface{}]interface{} // map
	Lock sync.RWMutex                // 读写锁
}

// NewMap 创建并发 map
func NewMap() *SyncMap {
	return &SyncMap{
		Map: map[interface{}]interface{}{},
	}
}

// Get 并发读
func (sm *SyncMap) Get(key interface{}) (interface{}, bool) {
	sm.Lock.RLock()
	defer sm.Lock.RUnlock()

	v, ok := sm.Map[key]
	return v, ok
}

// GetOrSet 获取不到就设置
func (sm *SyncMap) GetOrSet(key interface{}, value interface{}) interface{} {
	sm.Lock.RLock()

	// 获取 key
	if v, ok := sm.Map[key]; ok {
		sm.Lock.RUnlock()
		return v
	}
	sm.Lock.RUnlock()

	// 设置 key、value
	sm.Set(key, value)

	return value
}

// Set 并发写
func (sm *SyncMap) Set(key interface{}, value interface{}) {
	sm.Lock.Lock()
	defer sm.Lock.Unlock()

	sm.Map[key] = value
}

// Exists 并发判断
func (sm *SyncMap) Exists(key interface{}) bool {
	sm.Lock.RLock()
	defer sm.Lock.RUnlock()

	if _, ok := sm.Map[key]; ok {
		return true
	}

	return false
}

// Delete 并发删
func (sm *SyncMap) Delete(key interface{}) {
	sm.Lock.Lock()
	defer sm.Lock.Unlock()

	delete(sm.Map, key)
}

// Range 遍历
func (sm *SyncMap) Range(f func(key, value interface{})) {
	sm.Lock.RLock()
	defer sm.Lock.RUnlock()

	for key, value := range sm.Map {
		f(key, value)
	}
}
