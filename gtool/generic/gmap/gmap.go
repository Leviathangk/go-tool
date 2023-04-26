package gmap

import "sync"

// SyncMap 并发 map
type SyncMap[KEY comparable, VALUE any] struct {
	Map  map[KEY]VALUE // map
	Lock sync.RWMutex  // 读写锁
}

// NewMap 创建并发 map
func NewMap[KEY comparable, VALUE any]() *SyncMap[KEY, VALUE] {
	return &SyncMap[KEY, VALUE]{
		Map: map[KEY]VALUE{},
	}
}

// Get 并发读
func (sm *SyncMap[KEY, VALUE]) Get(key KEY) (VALUE, bool) {
	sm.Lock.RLock()
	defer sm.Lock.RUnlock()

	v, ok := sm.Map[key]
	return v, ok
}

// GetOrSet 获取不到就设置
func (sm *SyncMap[KEY, VALUE]) GetOrSet(key KEY, value VALUE) VALUE {
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
func (sm *SyncMap[KEY, VALUE]) Set(key KEY, value VALUE) {
	sm.Lock.Lock()
	defer sm.Lock.Unlock()

	sm.Map[key] = value
}

// Exists 并发判断
func (sm *SyncMap[KEY, VALUE]) Exists(key KEY) bool {
	sm.Lock.RLock()
	defer sm.Lock.RUnlock()

	if _, ok := sm.Map[key]; ok {
		return true
	}

	return false
}

// Delete 并发删
func (sm *SyncMap[KEY, VALUE]) Delete(key KEY) {
	sm.Lock.Lock()
	defer sm.Lock.Unlock()

	delete(sm.Map, key)
}

// Range 遍历
func (sm *SyncMap[KEY, VALUE]) Range(f func(key KEY, value VALUE)) {
	sm.Lock.RLock()
	defer sm.Lock.RUnlock()

	for key, value := range sm.Map {
		f(key, value)
	}
}
