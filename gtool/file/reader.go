// package file reader 是用来快速读取文件的
package file

import (
	"bufio"
	"io"
	"os"
)

// ReadAll 一次性读取文件
func ReadAll(path string) ([]byte, error) {
	return os.ReadFile(path)
}

// ReadLine 逐行读取文件
func ReadLine(path string, f func(line []byte)) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		f(scanner.Bytes())
	}

	return nil
}

// ReadByte 每次读取指定字节
func ReadByte(path string, size int, f func(b []byte)) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	buf := make([]byte, size)
	for {
		byteRead, err := file.Read(buf)
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}

		f(buf[:byteRead])
	}

	return nil
}
