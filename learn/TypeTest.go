package main

import (
	"fmt"
	"reflect"
)

// 自定义类型,拥有int的功能
type myInt int

// 别名
type intAlia = int

// 变相对int进行增加方法
func (m myInt) printSelf() {
	fmt.Println(m)
}

// 父结构体
type Person struct {
	name string
	age  int
	sex  string
}

// 子结构体
type Son struct {
	name   string
	Person // 匿名变量(其实就是相当于person Person,默认就是类型名称,注意:不能有2个一样类型的匿名变量) , 相当于继承Person
}

func (s Son) run() {
	fmt.Println(s.name)
}

// 这个函数只能Person调用,无法修改
func (p Person) printPerson() {
	fmt.Printf("姓名:%v,年龄:%v,性别:%v\n", p.name, p.age, p.sex)
}

// 因为传入的是指针,所以值可以被修改
func (p *Person) setInfo(name string, age int) {
	p.age = age
	p.name = name
}

func main() {

	// 自定义类型
	var a myInt = 10
	fmt.Println(a)
	// main.myInt
	fmt.Println(reflect.TypeOf(a))
	a.printSelf()

	// 别名
	var b intAlia = 80
	fmt.Println(b)
	// int
	fmt.Println(reflect.TypeOf(b))

	var p = Person{"DogLi", 45, "男"}
	p.printPerson()
	p.setInfo("大佬", 88)
	p.printPerson()

	// 测试继承结构体
	var s Son
	s.name = "Son"
	s.Person = Person{"Person", 11, "男"}
	// Son原本是不能调用printPerson(因为这是Person才可以调的),但是由于Son结构体中含有Person,所以当在Son找不到对应的方法时,会在Person中查找
	// person的方法
	s.printPerson()
	// son的方法
	s.run()
	// 直接给person的属性age赋值,如果对应的属性在son中找不到,那么会从结构体中其他结构体中查找
	s.age = 88
}
