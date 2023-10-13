package main

import (
	"fmt"
	"reflect"
)

type Hello struct {
	Name  string
	Hobby string
	skill string
}

func (h Hello) Say() {
	fmt.Println(h.Name)
}

// 反射测试
func main() {
	x := 10
	reflectPrint(x)
	var u = Hello{"念念", "睡觉", "coder"}
	reflectStruct(u)
}

func reflectPrint(x interface{}) {
	// 获取类型
	//t := reflect.TypeOf(x)
	var num int64 = 10
	value := reflect.ValueOf(&num)
	fmt.Println(num)
	value.Elem().SetInt(666)
	fmt.Println(num)
}

// 通过反射确定类型
func typeValue(v reflect.Value) {
	kind := v.Kind()
	switch kind {
	case reflect.Int64:
		fmt.Println("Int类型")
		break
	case reflect.Float64:
		fmt.Println("Float64类型")
		break
	case reflect.Float32:
		fmt.Println("Float32类型")
		break
	case reflect.String:
		fmt.Println("String类型")
		break
	default:
		fmt.Println("未知类型")
	}
}

// 通过反射设置值
func reflectSetValue() {
	var num int64 = 10
	value := reflect.ValueOf(&num)
	fmt.Println(num)
	value.Elem().SetInt(666)
	fmt.Println(num)
}

// 结构体反射
func reflectStruct(x interface{}) {
	typeOf := reflect.TypeOf(x)
	field0 := typeOf.Field(0)

	fmt.Println("字段名称:", field0.Name)
	fmt.Println("字段类型:", field0.Type)
	fmt.Println("字段Tag:", field0.Tag)

	// 获取该结构体字段数量
	numField := typeOf.NumField()
	fmt.Println("字段数量:", numField)

	// 获取属性对应值
	valueOf := reflect.ValueOf(x)
	hobbyValue := valueOf.FieldByName("Hobby")
	fmt.Println(hobbyValue)

	method, _ := typeOf.MethodByName("Say")
	fmt.Println(method)
	// 执行结构体的方法
	valueOf.MethodByName("Say").Call(nil)
}

func testReflect() {
	var user1 User
	user1.Name = "HaHa"
	user1.Age = 60
	user1.P1.Name = "HHHHH"
	fmt.Println(user1)

	typeOf := reflect.TypeOf(user1)
	field := typeOf.NumField()
	fmt.Println(field)
	fmt.Println(typeOf.NumMethod())

	structField := typeOf.Field(0)
	fmt.Println("===========================================")
	fmt.Println(structField.Name)
	fmt.Println(structField.Type)
	fmt.Println(structField.Index)
	fmt.Println("===========================================")

	valueOf := reflect.ValueOf(user1)
	name := valueOf.FieldByName("Name")
	fmt.Println(name)
	fmt.Println(valueOf.Field(3))

	// 反射执行无参函数
	valueOf.MethodByName("Say").Call(nil)
	// 反射执行有参函数
	values := valueOf.MethodByName("Add").Call([]reflect.Value{reflect.ValueOf(int32(1))})
	fmt.Println(values[0])

	// 通过反射获取值
	value := valueOf.FieldByName("Age").Int()
	fmt.Println(value + 1)

	// 通过反射 赋值
	reflect.ValueOf(&user1).Elem().FieldByName("Age").SetInt(88)

	fmt.Println(user1.Age)
}

type Person struct {
	Name   string
	Age    int32
	Adress string
}

type User struct {
	Name string
	Age  int32
	P1   Person
	P2   Person
	Person
}

func (u User) Say() {
	fmt.Println(u.Name)
}

func (u User) Add(x int32) int32 {
	return u.Age + x
}
