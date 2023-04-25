package file

import (
	"os"
	"path/filepath"
)

// MoveTo 移动指定文件到另一个位置
// 不区分文件文件夹
func MoveTo(oldpath, newpath string) error {
	parentPath := filepath.Dir(newpath)
	_, err := os.Stat(parentPath)
	if err != nil {
		err = os.MkdirAll(parentPath, 0666)
		if err != nil {
			return err
		}
	}
	return os.Rename(oldpath, newpath)
}
