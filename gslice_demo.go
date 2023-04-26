package main

import (
	"fmt"
	"github.com/Leviathangk/go-tool/gtool/generic/gslice"
)

func gsliceDemo() {
	res := gslice.Insert([]int{1, 2, 3, 4, 5, 6}, 0, 8)
	fmt.Println(res)

	res = gslice.Remove([]int{1, 1, 5, 67, 1, 2, 1, 3, 1}, 1, 0)
	fmt.Println(res)

	res2 := gslice.Exists([]int{1, 2, 3, 4, 5, 6}, 7)
	fmt.Println(res2)
}
