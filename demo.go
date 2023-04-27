package main

import (
	"fmt"
	"github.com/Leviathangk/go-tool/gtool/generic/gmap"
)

func main() {
	m := gmap.NewMap[string, string]()
	m.Set("key", "value")
	fmt.Println(m.Get("key"))

	m.Range(func(key string, value string) {
		fmt.Printf("%v -> %v\n", key, value)
	})

	fmt.Println(m.GetOrSet("key2", "value2"))
}
