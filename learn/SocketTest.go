package main

import (
	"fmt"
	"net"
	"time"
)

// socket 测试
//
func main() {
	//service := ":8686"
	//tcpAddr, _ := net.ResolveTCPAddr("tcp4", service)
	//listener, _ := net.ListenTCP("tcp", tcpAddr)
	listener, _ := net.Listen("tcp", "127.0.0.1:8888")
	for {
		conn, _ := listener.Accept()
		fmt.Println("建立连接")
		go readOrWrite(conn)
	}
}

func readOrWrite(conn net.Conn) {
	defer closeResource(conn)
	bytes := make([]byte, 4)
	for {
		//conn.Write([]byte("Hello"))
		//conn.SetReadDeadline(time.Now().Add(time.Second * 5))
		read, err := conn.Read(bytes)
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Printf("长度=%d,%s \n", read, time.Now())
		// 字节数组转字符串
		fmt.Println(string(bytes[0:read]))
	}
}

func closeResource(conn net.Conn) {
	fmt.Println("finally代码块->资源关闭")
	_ = conn.Close()
}
