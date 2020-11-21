package main

import (
	"errors"
	"fmt"
	"math/rand"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"
)

/**
  package main : main声明这个文件是一个可执行的,是可以直接运行的

  相当于Java的main函数
*/

func main() {
	//var i, s int
	//i = 10
	//s = 20
	//swap(i, s)
	//fmt.Println(i)
	//swapReference(&i, &s)
	//array()
	//point()
	//nullPoint()
	//arrayPoint()
	//pointPoint()
	//testStruct()
	//slice()
	//mapTest()
	//typeTrans()
	//testInterface()
	//err := testError(-1)
	//if err != nil {
	//	fmt.Println(err)
	//}

	// 并发(就是多线程了)
	//go concurrent("HelloWorld")
	//go concurrent("HelloWorld2")
	//time.Sleep(time.Second * 5)
	//channelTest()
	// defer修饰的函数: 在函数结束的时候自动调用! 在此处会在lock()结束后运行
	//defer func() {
	//	fmt.Println("处理异常开始")
	//	for err := recover(); err != nil; {
	//		fmt.Println(err)
	//	}
	//	fmt.Println("处理异常结束")
	//}()
	//lock()
	//stringTest()
}

// 返回多个值,并已经事先定义好了返回变量(默认值)
func returnMulti(x, y int) (i int, v int) {
	//i = x
	//v = y
	return
}

// 字符串常用
func stringTest() {
	var str string = "你好"
	for s := range str {
		// 这样打印的出来中文是乱码
		fmt.Println(string(s))
	}

	// 正确打印中文字符串
	runes := []rune(str)
	for i := 0; i < len(runes); i++ {
		fmt.Println(string(runes[i]))
	}

	str = "Hello-World"
	contains := strings.Contains(str, "He")
	fmt.Println(contains)
	// 切割
	split := strings.Split(str, "-")
	fmt.Println(split)
	// 合并字符串数据,指定分隔符合并
	join := strings.Join(split, "*")
	fmt.Println(join)
	// 比较字符串(不分大小写)
	fold := strings.EqualFold("Hello", "hello")
	fmt.Println(fold) // true
	// 统计子串出现了多少次
	count := strings.Count("Hellolllo", "l")
	fmt.Println(count)
	// 指定替换某个元素,替换数量为1
	replace := strings.Replace("Hello", "l", "F", 1)
	fmt.Println(replace)
	// 转为小写
	lower := strings.ToLower("Hello")
	fmt.Println(lower)

	// 数字转字符串,指定进制
	formatInt := strconv.FormatInt(4444444, 10)
	fmt.Println(reflect.TypeOf(formatInt))
	fmt.Println(formatInt)

	// 字符串转数字
	parseInt, _ := strconv.ParseInt("1", 10, 64)
	fmt.Println(parseInt)
	fmt.Println(reflect.TypeOf(parseInt))
}

// 锁 -> 互拆锁和读写锁
func lock() {
	// 创建一个锁对象
	var x int = 0
	var y int = 10
	var mutex sync.Mutex
	for i := 0; i < 10000; i++ {
		go func() {
			mutex.Lock()
			x = x / y
			y++
			mutex.Unlock()
		}()
	}
	time.Sleep(time.Second)
	fmt.Println(x)
}

// 通道使用
func channelTest() {
	//通道
	ch := make(chan int64)
	// ch := make(chan int,100) 设置缓冲区的大小为100
	go func() {
		for i := 0; i < 5; i++ {
			result := rand.Int63n(time.Now().Unix())
			fmt.Println(result)
			// 把这个num值放在 c 通道中
			ch <- result
		}
		// 关闭通道
		close(ch)
	}()

	// 从通道取出数据一
	//x,y := <-ch,<- ch
	//fmt.Println(x,y)
	//===========================================
	//通道取出数据二
	//fmt.Println("通道取出数据二:",<-ch,<-ch)
	// 使用range遍历通道数据,必须使用这种方式,否则会报死锁
	//go func() {
	//	for k := range ch {
	//		fmt.Println(k)
	//	}
	//}()
	// time.Sleep(time.Second * 5)
	//===========================================
	for k := range ch {
		fmt.Println(k)
	}
}

// 并发
func concurrent(s string) {

	time.Sleep(time.Millisecond * 1000)
	fmt.Println(s)
}

//异常
func testError(num int) error {
	if num < 0 {
		return errors.New("参数不能为负数")
	}
	return nil
}

// 接口测试
func testInterface() {
	// 多态
	var work Work
	work = Coder{name: "程序员", age: 25}
	work.work()
	work.say()

	work = Teacher{name: "教师", age: 50}
	work.work()
	work.say()
}

// 定义一个接口
type Work interface {
	work()
	say() string
}

// 程序员
type Coder struct {
	name string
	age  int
}

// 接口实现一
func (coder Coder) work() {
	fmt.Println("我上班写代码")
}
func (coder Coder) say() string {
	fmt.Println("我996")
	return coder.name
}

// 教师
type Teacher struct {
	name string
	age  int
}

// 接口实现二
func (teacher Teacher) work() {
	fmt.Println("我上班教书")
}

func (teacher Teacher) say() string {
	fmt.Println("我865")
	return teacher.name
}

// 类型转换
func typeTrans() {
	var a int = 10
	var b float32 = float32(a)
	fmt.Println(b)
	// 使用反射来进行类型推断
	fmt.Println(reflect.TypeOf(b))
}

// map
func mapTest() {
	// map定义的格式
	var testMap map[string]string
	// 初始化
	testMap = make(map[string]string)
	// put数据
	testMap["name"] = "HelloWorld"
	testMap["age"] = "19"
	fmt.Println(testMap)
	// 获取数据
	fmt.Println(testMap["age"])
	fmt.Println("================================")

	test2Map := map[string]string{"name": "six", "age": "18"}
	// 删除元素
	delete(test2Map, "name")
	fmt.Println(test2Map)

	test2Map["name"] = "Hello"
	test2Map["address"] = "beijing"
	fmt.Printf("%v", len(test2Map))
}

// 切片
func slice() {
	var nums [5]int = [...]int{1, 2, 3, 4, 5}
	// 引用nums数值的索引1到4的值(不含4)
	slice := nums[1:4]
	fmt.Println("slice长度:", len(slice))
	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}
	// 改变了slice值,但是不改变nums中的值
	slice[0] = 10
	fmt.Println(nums[0])
	fmt.Println(slice[0])
	fmt.Println("-----------------")
	var r *int = &slice[0]
	fmt.Println(r)
	fmt.Println(&nums[0])
	// range遍历 i是索引,v是值
	for i, v := range nums {
		fmt.Printf("i=%v,v=%v\n", i, v)
	}
	fmt.Println("-----------------")
	slice2 := make([]int, 5, 10)
	fmt.Println(slice2)
	for i := 0; i < len(slice2); i++ {
		slice2[i] = i + 10
	}
	fmt.Println(slice2)
	fmt.Println("-----------------")

	var iArray []int
	iArray = append(iArray, 0)
	iArray = append(iArray, 1)
	iArray = append(iArray, 2)
	fmt.Println(iArray)
	// 创建切片
	bArray := make([]int, len(iArray), cap(iArray)*2)
	// 数组拷贝,将数组iArray数据拷贝到bArray
	copy(bArray, iArray)
	fmt.Println(bArray)
}

type Man struct {
	name    string
	age     int
	sex     string
	address string
}

func testStruct() {
	// 创建方式一
	man := Man{"TT", 1, "男", "北京"}
	// 创建方式二
	man2 := Man{name: "NN", age: 18, address: "", sex: "男"}
	fmt.Println(man)
	fmt.Println(man2)

	var man3 Man
	man3.sex = "男"
	man3.address = "北京市"
	man3.age = 77

	fmt.Println(man3)
	fmt.Println(man3.age)

	printStruct(man3)
	fmt.Println(man3.age)
}

func printStruct(man Man) {
	fmt.Println(man.age)
	fmt.Println(man.sex)
	// 不会改变外面的数值,也就是说只是值传递
	man.age = 666
	fmt.Println(man.age)
}

func printStruct2(man *Man) {
	fmt.Println(man.age)
	fmt.Println(man.sex)
	//会改变外面的值,因为这是 指针
	man.age = 666
	fmt.Println(man.age)
}

//指针的指针
func pointPoint() {
	var a int = 10
	var b *int
	var c **int
	b = &a
	c = &b
	fmt.Println("a=", &a)
	fmt.Println("b=", b)
	fmt.Println("*b=", *b)
	fmt.Println("c地址=", c)
	fmt.Println("*c=", *c)
	fmt.Println("**c=", **c)
}

// 数组指针
// 数组的每一个元素其实都是指针
func arrayPoint() {
	var i = []int{10, 20, 30}
	var ptr [3]*int
	for k := 0; k < len(i); k++ {
		ptr[k] = &i[k]
	}
	var num int = 66
	var ptrNum *int = &num
	ptr[0] = ptrNum
	fmt.Println(*ptr[0])
	fmt.Println(i[0])

	ptr[0] = &i[0]
	i2 := ptr[0]
	*i2 = 88
	// 打印的是值
	fmt.Println("*i2=", *i2)
	// 打印的是地址
	fmt.Println("i2=", i2)
	fmt.Println(*ptr[0])
	fmt.Println(i[0])
}

// 数组
func array() {
	var array = [10]int{1, 2, 3, 4}
	fmt.Println(array[0])
	// 不设置数组的大小,那么会根据元素的大小进行设置
	var str = [...]string{"H", "I", "J"}
	// 数组遍历
	for i := 0; i < len(str); i++ {
		fmt.Println(str[i])
	}

	// 这个是切片
	k := []int{6, 3, 9, 10, 1, 2, 3, 4}
	fmt.Println(k[0])
	// 对数组进行排序
	sort.Ints(k)
	fmt.Println(k)
}

// 空指针
func nullPoint() {
	var ptr *int
	if ptr != nil {
		fmt.Println("非空")
	} else {
		fmt.Println("空指针")
	}
	// 打印nil
	fmt.Println(ptr)
}

// 打印指针地址
//一个指针变量指向了一个值的内存地址。
func point() {
	var i int = 10
	var j *int
	j = &i
	// 打印指针地址
	fmt.Println(&i)
	// 打印指针地址
	fmt.Println(j)
	// 改变指针指向的值,同时也改变了i值,因为i和j都是指向同一个地址
	*j = 30
	// 打印指针指向的实际值(非地址)
	fmt.Println(*j)
	fmt.Println(i)
}

func swap(i, s int) {
	var temp = i
	i = s
	s = temp
}

// 指针
func swapReference(i *int, s *int) {
	var temp int
	temp = *i
	*i = *s
	*s = temp

}
