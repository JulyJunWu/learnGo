package main

import (
	"fmt"
	"reflect"
)

// 自定义函数
func main() {
	var h hello
	// 函数赋值
	h = add
	fmt.Println(reflect.TypeOf(h))
	result := h(1, 1)
	fmt.Println(result)
	h = sub
	result = h(1, 1)
	fmt.Println(result)

	// 将自定义函数作为参数
	result = calc(10, 20, h)
	fmt.Println(result)

	// 匿名函数,直接执行
	func() {
		fmt.Println("匿名函数")
	}()

	// 匿名函数赋值
	var f = func(x, y int) {
		fmt.Println(x - y)
	}
	// 执行匿名函数
	f(1, 10)
}

// 自定义类型函数 , 只有满足该参数和返回类型的 函数才可以赋值
type hello func(x, y int) int

func add(x, y int) int {
	return x + y
}

func sub(x, y int) int {
	return x - y
}

// 函数作为参数
func calc(x, y int, f hello) int {
	return f(x, y)
}

// 函数作为返回值
func returnFunc(t int) hello {
	switch t {
	case 1:
		return sub
	case 2:
		return add
	default:
		return nil
	}
}
