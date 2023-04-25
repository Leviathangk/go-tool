// Package gsync 并发 map，但是获取数据后需要转换
package gsync

import "sync"

// SyncMap 并发 map
type SyncMap struct {
	Map  *map[interface{}]interface{} // map
	lock sync.RWMutex                 // 读写锁
}

// NewMap 创建并发 map
func NewMap() *SyncMap {
	return &SyncMap{
		Map: &map[interface{}]interface{}{},
	}
}

// Get 并发读
func (sm *SyncMap) Get(key interface{}) interface{} {
	sm.lock.RLock()
	defer sm.lock.RUnlock()

	m := *sm.Map
	if v, ok := m[key]; ok {
		return v
	}
	return nil
}

// Set 并发写
func (sm *SyncMap) Set(key interface{}, value interface{}) {
	sm.lock.Lock()
	defer sm.lock.Unlock()

	m := *sm.Map
	m[key] = value
}

// Exists 并发判断
func (sm *SyncMap) Exists(key interface{}) bool {
	sm.lock.RLock()
	defer sm.lock.RUnlock()

	m := *sm.Map
	if _, ok := m[key]; ok {
		return true
	}

	return false
}

// Delete 并发删
func (sm *SyncMap) Delete(key interface{}) {
	sm.lock.Lock()
	defer sm.lock.Unlock()

	m := *sm.Map
	delete(m, key)
}
