package main

import (
	"fmt"
	"github.com/Leviathangk/go-tool/gtool/generic/gmap"
	"sync"
)

func mapDemo() {
	m := gmap.NewMap[string, string]()
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
