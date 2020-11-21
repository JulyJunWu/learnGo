package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"reflect"
	"strconv"
)

// 启动http服务,监听端口
//
func main() {
	http.HandleFunc("/", handle)
	go startHttp(6666)
	go startHttp(8888)
	http.ListenAndServe(fmt.Sprintf("127.0.0.1:%d", 9999), nil)
}

func startHttp(port int) {
	http.ListenAndServe(fmt.Sprintf("127.0.0.1:%d", port), nil)
}

type Msg struct {
	Code int `json:"code"`
	// 使用interface就可以接收不同的类型,不必局限与单一类型
	Data map[string]interface{} `json:"data"`
	// 存放任意类型
	Msg interface{} `json:"msg"`
}

func handle(w http.ResponseWriter, r *http.Request) {
	//urlStr := r.URL.String()
	// 请求路径
	path := r.URL.Path
	//fmt.Println(urlStr)
	//fmt.Print(r.URL.Path)
	//fmt.Println(r.Method)
	//fmt.Println(r.Host)
	//fmt.Println(r.Proto)
	//fmt.Println(r)
	// 获取参数,就是把参数变为Map的形式
	params := r.URL.Query()
	fmt.Println("参数:", params)
	strAge := params["age"]
	var age string
	if strAge != nil {
		age = strAge[0]
	}
	intAge, err := strconv.Atoi(age)
	if err != nil {
		fmt.Println(err)
	} else {
		// 打印类型
		fmt.Println(reflect.TypeOf(intAge))
	}
	fmt.Println("host:", r.Host)
	switch path {
	case "/hello":
		io.WriteString(w, "你好,世界!")
	case "/nice":
		io.WriteString(w, "非常漂亮!")
	case "/json":
		// 返回json数据
		w.Header().Set("content-type", "text/json")
		mapData := make(map[string]interface{})
		mapData["姓名"] = "Hello"
		mapData["年龄"] = 18
		mapData["地址"] = "北京市"
		msg := Msg{Code: 200, Data: mapData, Msg: 666}
		// 解析为json数据
		data, _ := json.Marshal(msg)
		w.Write(data)
	default:
		io.WriteString(w, "欢迎来到首页!")
	}
	//var strByte = []byte("欢迎!")
	//w.Write(strByte)
}
