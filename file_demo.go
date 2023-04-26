package main

import (
	"fmt"
	"github.com/Leviathangk/go-tool/gtool/file"
)

func fileDemo() {
	// content, err := file.ReadAll("D:\\ddd.txt")
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// fmt.Println(string(content))

	// err := file.ReadLine("D:\\ddd.txt", func(line []byte) {
	// 	fmt.Println(string(line))
	// })
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// err := file.ReadByte("D:\\ddd.txt", 1024, func(b []byte) {
	// 	fmt.Println(string(b))
	// })
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// writer, err := file.WriteByte("D:\\ddd2.txt", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// defer writer.Close()

	// _ = writer.Write([]byte("1"))
	// _ = writer.Write([]byte("2"))
	// _ = writer.Write([]byte("3"))

	err := file.MoveTo("D:\\dddd", "D:\\dddd2")
	fmt.Println(err)
}
