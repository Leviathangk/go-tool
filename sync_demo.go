package main

import (
	"fmt"
	"github.com/Leviathangk/go-tool/gtool/gsync"
	"sync"
)

func syncDemo() {
	// 并发 slice
	s := gsync.NewSlice[int]()
	s.Append(1, 2, 3, 4)
	fmt.Println(s.Slice)

	co := s.Copy()

	s.Range(func(index int, value int) {
		fmt.Printf("%d -> %d\n", index, value)
	})

	fmt.Println(s.IndexOf(4))

	fmt.Println(s.Exists(5))

	fmt.Println(s.Get(-2))

	s.Insert(0, 0)
	fmt.Println(s.Slice)

	s.Remove(4, 0)
	fmt.Println(s.Slice)

	s.RemoveByIndex(s.Length() - 1)
	fmt.Println(s.Slice)

	p := s.Pop()
	fmt.Println(p)
	fmt.Println(s.Slice)

	sh := s.Shift()
	fmt.Println(sh)
	fmt.Println(s.Slice)

	fmt.Println(co)

	// 并发 map
	m := gsync.NewMap[string, string]()
	m.Set("key", "value")
	fmt.Println(m.Get("key"))

	m.Range(func(key string, value string) {
		fmt.Printf("%v -> %v\n", key, value)
	})

	fmt.Println(m.GetOrSet("key2", "value2"))

	// 自带的并发 map
	var ma sync.Map
	ma.Store("key", "value")
	fmt.Println(ma.Load("key"))

	ma.Range(func(key, value any) bool {
		fmt.Printf("%v -> %v\n", key, value)
		return true
	})
}
