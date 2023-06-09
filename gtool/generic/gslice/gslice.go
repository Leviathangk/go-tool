package gslice

// Append 切片添加元素
func Append[T any](s []T, elems ...T) []T {
	return append(s, elems...)
}

// Insert 在指定索引插入数据
func Insert[T any](s []T, index int, elems T) []T {
	return append(s[:index], append([]T{elems}, s[index:]...)...)
}

// Remove 移除指定值的数据
// nums <= 0 清除所有指定值的数据
// nums > 0 清除指定数量的指定值的数据
func Remove[T comparable](s []T, value T, nums int) []T {
	count := 0
	length := len(s)

	for i := 0; i < length; i++ {
		if s[i-count] == value {
			if nums <= 0 {
				s = append(s[:i-count], s[i-count+1:]...)
				count += 1
			} else if nums > 0 && count != nums {
				s = append(s[:i-count], s[i-count+1:]...)
				count += 1
			} else {
				break
			}
		}
	}

	return s
}

// RemoveByIndex 移除指定索引的数据
func RemoveByIndex[T any](s []T, index int) []T {
	return append(s[:index], s[index+1:]...)
}

// Pop 移除末尾元素
func Pop[T any](s []T) ([]T, T) {
	index := len(s) - 1
	last := s[index]

	return RemoveByIndex(s, index), last
}

// Shift 移除头部元素
func Shift[T any](s []T) ([]T, T) {
	first := s[0]

	return RemoveByIndex(s, 0), first
}

// Copy 复制一份，返回副本
func Copy[T any](s []T) []T {
	newS := make([]T, len(s))
	copy(newS, s)

	return newS
}

// Range 遍历
func Range[T any](s []T, f func(index int, value T)) {
	for index := range s {
		f(index, s[index])
	}
}

// IndexOf 返回指定值第一次出现的索引，无则为 -1
func IndexOf[T comparable](s []T, v T) int {
	for index, value := range s {
		if value == v {
			return index
		}
	}

	return -1
}

// Exists 判断指定值是否存在
func Exists[T comparable](s []T, v T) bool {
	return IndexOf(s, v) != -1
}
