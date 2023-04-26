package main

import (
	"fmt"
	"github.com/Leviathangk/go-tool/gtool/gsync"
)

func main() {
	// 并发 slice
	s := gsync.NewSlice[int]()
	s.Append(1, 2, 3, 4)

	res := s.Shift()
	fmt.Println(res)
	fmt.Println(s.Slice)

	// 并发 map
	m := gsync.NewMap[int, int]()
	m.Set(1, 1)
	m.Set(2, 2)

	v, ok := m.Get(1)
	if ok {
		fmt.Println(v)
	}
}
