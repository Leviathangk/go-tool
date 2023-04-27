package gslice

// GSlice 泛型切片
type GSlice[T comparable] struct {
	Slice []T // GSlice
}

// NewSlice 创建切片
func NewSlice[T comparable]() *GSlice[T] {
	return &GSlice[T]{
		Slice: []T{},
	}
}

// Append 添加元素
func (gs *GSlice[T]) Append(ele ...T) {
	gs.Slice = append(gs.Slice, ele...)
}

// Insert 在指定索引插入数据
func (gs *GSlice[T]) Insert(index int, elems T) {
	gs.Slice = append(gs.Slice[:index], append([]T{elems}, gs.Slice[index:]...)...)
}

// Remove 移除指定值的数据
// nums <= 0 清除所有指定值的数据
// nums > 0 清除指定数量的指定值的数据
func (gs *GSlice[T]) Remove(value T, nums int) {
	count := 0
	length := len(gs.Slice)

	for i := 0; i < length; i++ {
		if gs.Slice[i-count] == value {
			if nums <= 0 {
				gs.Slice = append(gs.Slice[:i-count], gs.Slice[i-count+1:]...)
				count += 1
			} else if nums > 0 && count != nums {
				gs.Slice = append(gs.Slice[:i-count], gs.Slice[i-count+1:]...)
				count += 1
			} else {
				break
			}
		}
	}
}

// RemoveByIndex 移除指定索引的数据
func (gs *GSlice[T]) RemoveByIndex(index int) {
	gs.Slice = append(gs.Slice[:index], gs.Slice[index+1:]...)
}

// Pop 移除末尾元素
func (gs *GSlice[T]) Pop() T {
	index := len(gs.Slice) - 1
	last := gs.Slice[index]
	gs.Slice = append(gs.Slice[:index], gs.Slice[index+1:]...)

	return last
}

// Shift 移除头部元素
func (gs *GSlice[T]) Shift() T {
	first := gs.Slice[0]
	gs.Slice = append([]T{}, gs.Slice[1:]...)

	return first
}

// Copy 复制一份，返回副本
func (gs *GSlice[T]) Copy() []T {
	newS := make([]T, gs.Length())
	copy(newS, gs.Slice)

	return newS
}

// Range 遍历
func (gs *GSlice[T]) Range(f func(index int, value T)) {
	for index := range gs.Slice {
		f(index, gs.Slice[index])
	}
}

// IndexOf 返回指定值第一次出现的索引，无则为 -1
func (gs *GSlice[T]) IndexOf(v T) int {
	for index, value := range gs.Slice {
		if value == v {
			return index
		}
	}

	return -1
}

// Exists 判断指定值是否存在
func (gs *GSlice[T]) Exists(v T) bool {
	return gs.IndexOf(v) != -1
}

// Get 获取元素：支持负数访问
func (gs *GSlice[T]) Get(index int) T {
	if index < 0 {
		index = gs.Length() + index
	}

	return gs.Slice[index]
}

// Length 获取长度
func (gs *GSlice[T]) Length() int {
	return len(gs.Slice)
}
