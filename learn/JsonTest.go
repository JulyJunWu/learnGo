package main

import (
	"encoding/json"
	"fmt"
)

// json使用测试
func main() {

	var s = Student{"小三", 188, "男", "上海"}
	// 把结构体变为json的byte数组切片
	jsonByte, err := json.Marshal(s)
	if err != nil {
		fmt.Println(err)
	}
	// 转为json字符串
	jsonStr := string(jsonByte)
	fmt.Println(jsonStr)

	// 把json字符串数据转为 结构体
	var str = `{"Name":"大表哥","Age":18,"Sex":"男","Address":"中国"}`
	var student Student
	// 先要转为字节切片数组
	_ = json.Unmarshal([]byte(str), &student)
	fmt.Println(student.Name)

	//--------------嵌套
	var class = Class{"高三6班", make([]Student, 0, 6)}
	for i := 0; i <= 30; i++ {
		var st = Student{fmt.Sprintf("张三%d", i), 12, "男", "shanghai"}
		class.Student = append(class.Student, st)
	}
	selfPrint(class)
	byteClass, _ := json.Marshal(class)
	selfPrint(string(byteClass))
	// 转为结构体
	var class2 Class
	_ = json.Unmarshal(byteClass, &class2)
	//var class3 = &Class{}
	//_ = json.Unmarshal(byteClass, class3)
	selfPrint(class2)
}

// 如果要转为json的话,变量名一定要大写(因为要对外访问,大写为public,小写为private)
type Student struct {
	Name    string `json:"name"` // 这样json字符串就是小写的
	Age     int
	Sex     string
	Address string
}

type Class struct {
	Title   string
	Student []Student
}

func selfPrint(all interface{}) {
	fmt.Println(all)
}
