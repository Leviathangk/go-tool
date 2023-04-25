// package file reader 是用来快速读取文件的
package file

import (
	"os"
)

// WriteAll 一次性写文件
// path 文件路径含名字
// data 写入的数据
// flag 打开方式
// perm 权限
func WriteAll(path string, data []byte, flag int, perm os.FileMode) error {
	file, err := os.OpenFile(path, flag, perm)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(data)

	return err
}

type WriteByter struct {
	File *os.File
}

// WriteByte 逐字节写入（针对大文件）
// path 文件路径含名字
// flag 打开方式
// perm 权限
func WriteByte(path string, flag int, perm os.FileMode) (*WriteByter, error) {
	file, err := os.OpenFile(path, flag, perm)
	if err != nil {
		return nil, err
	}

	return &WriteByter{File: file}, nil
}

// Write 用来逐字节写入
func (w *WriteByter) Write(data []byte) error {
	_, err := w.File.Write(data)
	if err != nil {
		return err
	}

	err = w.File.Sync() // 刷新到硬盘
	if err != nil {
		return err
	}

	return nil
}

// Close 关闭文件
func (w *WriteByter) Close() error {
	return w.File.Close()
}
