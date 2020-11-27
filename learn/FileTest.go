package main

import (
	"fmt"
	"os"
)

// 文件测试
func main() {
	openFile, err := os.OpenFile("./test.txt", os.O_CREATE|os.O_WRONLY|os.O_RDONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer openFile.Close()
	_, _ = openFile.Write([]byte("耗子尾汁 \r\n"))
}
