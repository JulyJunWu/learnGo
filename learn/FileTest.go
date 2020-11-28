package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// 文件测试
func main() {

	//create()
	//truncate()
	//stat()
	//copy()
	//bufferWrite()
	read()
}

// 打开文件
func openFile() {
	openFile, err := os.OpenFile("./test.txt", os.O_CREATE|os.O_WRONLY|os.O_RDONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer openFile.Close()
	_, _ = openFile.Write([]byte("耗子尾汁 \r\n"))
}

// 创建文件
func create() {
	//file, err := os.Create("./hello.txt")
	file, err := os.Create("C:\\Users\\DELL\\Desktop\\资料\\ws.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	_, _ = file.Write([]byte("年轻人要讲码德!\n码林要以和为贵!"))
}

// 将文件保留多少字节
func truncate() {
	// 如果size为0,那么将文件里的内容全部清空!
	// 不为0(假设为N),则意味着将文件保留 N 个字节
	err := os.Truncate("hello.txt", 0)
	if err != nil {
		fmt.Println(err)
	}
	file, err := os.OpenFile("hello.txt", os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	_, _ = file.Write([]byte("来!骗!我,69岁的老码农!这好么?这不好!"))
}

//获得文件的信息(应该是元数据吧)
func stat() {
	file, _ := os.Stat("hello.txt")
	fmt.Println("Name:", file.Name())
	fmt.Println("IsDir:", file.IsDir())
	fmt.Println("Mode:", file.Mode())
	fmt.Println("Size:", file.Size())
	fmt.Println("ModTime:", file.ModTime())
	fmt.Println("Sys:", file.Sys())
}

// 复制文件
func copy() {
	srcFile, _ := os.Open("hello.txt")
	defer srcFile.Close()
	targetFile, _ := os.Create("hello2.txt")
	defer targetFile.Close()
	written, _ := io.Copy(targetFile, srcFile)
	fmt.Println(written)

	// 创建目录
	err := os.Mkdir("./ws", 0666)
	fmt.Println(err)
}

// 缓冲写
func bufferWrite() {
	file, _ := os.OpenFile("hello.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)
	defer file.Close()
	// 默认缓冲大小为4096字节(4K)
	writer := bufio.NewWriter(file)
	writeLen, _ := writer.WriteString("武林要以和为贵!")
	fmt.Println(writeLen)
	size := writer.Size()
	fmt.Println(size)
	available := writer.Available()
	fmt.Println(available)
	// 必须要刷新缓冲,将缓冲刷到磁盘
	_ = writer.Flush()
}

// 读取文件
func read() {
	file, _ := os.Open("hello.txt")
	defer file.Close()
	reader := bufio.NewReader(file)
	// 读缓冲大小
	size := reader.Size()
	for {
		// 读取一行
		line, prefix, err := reader.ReadLine()
		if err == io.EOF {
			fmt.Println("读取结束!")
			break
		}
		fmt.Println(err)
		fmt.Println(size)
		fmt.Println(len(line))
		fmt.Println(prefix)
		fmt.Println(string(line))
	}
}
