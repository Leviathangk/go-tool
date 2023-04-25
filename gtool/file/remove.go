package file

import "os"

// RemoveFile 删除文件
func RemoveFile(path string) error {
	return os.Remove(path)
}

// RemoveDir 删除文件夹
func RemoveDir(path string) error {
	return os.RemoveAll(path)
}
