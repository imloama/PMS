/*
 * 一些文件相关的操作，生成不同的内容
 */
package util

import (
	"fmt"
	"os"
)

func ReadHtml(filepath string) (content []byte, count int, err error) {
	file, err := os.Open(filepath)
	defer file.Close()
	if err != nil {
		fmt.Println("read file error:", filepath)
		return nil, 0, err
	}
	data := make([]byte, 1024)
	count, err2 := file.Read(data)
	if err2 != nil {
		fmt.Println("read file data error:", filepath)
		return nil, 0, err2
	}
	fmt.Printf("read %d bytes: %q\n", count, data[:count])

	return data, count, nil
}
