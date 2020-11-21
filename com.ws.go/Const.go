package main

import (
	"fmt"
	"unsafe"
)

/**
	常量
常量是一个简单值的标识符，在程序运行时，不会被修改的量。
常量中的数据类型只可以是布尔型、数字型（整数型、浮点型和复数）和字符串型。
常量的定义格式：
	const identifier [type] = value
*/

//	枚举
const (
	Unknown = 0
	Female  = 1
	Male    = 2
)

//	常量可以用len(), cap(), unsafe.Sizeof()函数计算表达式的值。
//	常量表达式中，函数必须是内置函数，否则编译不过：
const (
	Bad      = "bad"
	Good     = len(Bad)
	VeryGood = unsafe.Sizeof(Good)
)

/**
iota，特殊常量，可以认为是一个可以被编译器修改的常量。
iota 在 const关键字出现时将被重置为 0(const 内部的第一行之前)，
const 中每新增一行常量声明将使 iota 计数一次(iota 可理解为 const 语句块中的行索引)。
*/
const (
	A = iota // 0
	B        // 1
	C        // 2
	D = "HA" // HA iota = 2 + 1
	E        // HA iota = 3 + 1
	F = 100  // 100 iota = 4 + 1
	G        // 100 iota = 5 + 1
	H = iota // 7
	I        // 8
)

func main() {
	//	声明的常量可以不使用
	//	声明的变量必须使用
	const LENGTH int = 10
	const WIDTH int = 10
	var area int
	const a, b, c = 1, false, "str"
	area = LENGTH * WIDTH
	fmt.Printf("面积=%d", area)
	fmt.Println()
	fmt.Println(Unknown)

	fmt.Println(VeryGood)

	fmt.Println(D)
	fmt.Println(E)
	fmt.Println(F)
	fmt.Println(G)
	fmt.Println(H)
	fmt.Println(I)

}
