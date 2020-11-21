package main

import "fmt"

// defer测试
func main() {
	x := 1
	y := 2
	defer addNum(x, y, "one")
	x = 10
	defer addNum(x, y, "two")
	y = 20
	defer addNum(x, y, "three")

	throwException()
	fmt.Println("game over")
}

func addNum(x, y int, name string) int {
	result := x + y
	fmt.Printf("name:%v,x=%v,y=%v,result=%v \n", name, x, y, result)
	return result
}

func throwException() {
	defer func() {
		// 接收异常
		err := recover()
		if err != nil {
			fmt.Println("处理异常:", err)
		}
	}()
	// 如果抛出的异常未处理则程序退出!
	panic("throw 异常")
}
