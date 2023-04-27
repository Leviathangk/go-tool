package gmap

// GMap map
type GMap[KEY comparable, VALUE any] struct {
	Map map[KEY]VALUE
}

// NewMap 创建 map
func NewMap[KEY comparable, VALUE any]() *GMap[KEY, VALUE] {
	return &GMap[KEY, VALUE]{
		Map: map[KEY]VALUE{},
	}
}

// Get 读
func (sm *GMap[KEY, VALUE]) Get(key KEY) (VALUE, bool) {
	v, ok := sm.Map[key]
	return v, ok
}

// GetOrSet 获取不到就设置
func (sm *GMap[KEY, VALUE]) GetOrSet(key KEY, value VALUE) VALUE {
	// 获取 key
	if v, ok := sm.Map[key]; ok {
		return v
	}

	// 设置 key、value
	sm.Set(key, value)

	return value
}

// Set 写
func (sm *GMap[KEY, VALUE]) Set(key KEY, value VALUE) {
	sm.Map[key] = value
}

// Exists 判断
func (sm *GMap[KEY, VALUE]) Exists(key KEY) bool {
	if _, ok := sm.Map[key]; ok {
		return true
	}

	return false
}

// Delete 删
func (sm *GMap[KEY, VALUE]) Delete(key KEY) {
	delete(sm.Map, key)
}

// Range 遍历
func (sm *GMap[KEY, VALUE]) Range(f func(key KEY, value VALUE)) {
	for key, value := range sm.Map {
		f(key, value)
	}
}
