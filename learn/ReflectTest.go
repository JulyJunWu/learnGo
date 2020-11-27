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
