package main

import "fmt"

// 断言
func main() {
	// 声明一个空接口
	var a interface{} // 可以代表任何值
	num := 10
	a = num
	// 类型断言
	s, ok := a.(int)
	if ok {
		fmt.Println("类型是字符串")
	} else {
		fmt.Println("不是字符串")
	}
	fmt.Println(s)
	PrintType(s)

	var data = make(map[string]interface{})
	data["title"] = "超时"
	data["shop"] = Shop{"牙膏", 9.9}
	fmt.Println(data)
	shop := data["shop"]
	fmt.Println(shop)
	// 要想访问具体某个值,只能这样做
	v, ok := shop.(Shop)
	fmt.Println(v.Name)
}

type Shop struct {
	Name  string
	Price float64
}

// 通过switch推断
func PrintType(x interface{}) {
	switch x.(type) {
	case int:
		fmt.Println("int")
		break
	case string:
		fmt.Println("string")
		break
	default:
		fmt.Println("其他")
	}
}
