package main

import (
	"fmt"
	"sync"
	"time"
)

// 协程计数器
var group sync.WaitGroup

// 协程 并发
func main() {
	// 计数器加一
	group.Add(1)
	go concurrentTest()
	for i := 0; i < 10; i++ {
		fmt.Println("main", i)
		time.Sleep(time.Millisecond * 50)
	}
	// 等待计数器为0,主线程阻塞等待,直到所有的协程执行完毕
	group.Wait()
	fmt.Println("done")
}

func concurrentTest() {
	for i := 0; i < 10; i++ {
		fmt.Println("concurrentTest_", i)
		time.Sleep(time.Millisecond * 100)
	}
	// 计数器减一
	group.Done()
}
