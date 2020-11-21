package main

import (
	"fmt"
	"reflect"
	"time"
)

// 时间
func main() {
	now := time.Now()
	fmt.Println(now)
	// 日
	day := now.Day()
	fmt.Println(day)
	// 年
	year := now.Year()
	fmt.Println(year)
	// 月
	month := now.Month()
	fmt.Println(month)
	// 年月日
	date, m, d := now.Date()
	fmt.Println(date, m, d)
	// 时
	hour := now.Hour()
	fmt.Println(hour)
	// 分
	minute := now.Minute()
	fmt.Println(minute)
	// 秒
	second := now.Second()
	fmt.Println(second)
	// 纳秒
	nanosecond := now.Nanosecond()
	fmt.Println(nanosecond)
	local := now.Local()
	fmt.Println(local)
	// 格式化时间
	fmt.Println(now.Format("2006-01-02 15:04:05"))
	// 从1970开始的秒,时间戳
	fmt.Println(now.Unix())
	// 时间戳转time 1.秒 2.纳秒
	unix := time.Unix(1605951996, 0)
	fmt.Println(reflect.TypeOf(unix), unix)
	fmt.Println(unix.Format("2006-01-02 15:04:05"))
	//日期字符串转时间
	var tSr = "2020-11-21 17:46:36"
	var formatType = "2006-01-02 15:04:05"
	inLocation, _ := time.ParseInLocation(formatType, tSr, time.Local)
	fmt.Println(inLocation.Unix())

	// 加一小时
	t := inLocation.Add(time.Hour)
	fmt.Println(t)
	// 加一分钟
	t = inLocation.Add(time.Minute)
	fmt.Println(t)

	timer()
}

// 定时器
func timer() {
	// 一秒执行一次
	ticker := time.NewTicker(time.Second)
	i := 5
	for m := range ticker.C {
		i--
		if i < 0 {
			// 终止定时器
			ticker.Stop()
			break
		}
		fmt.Println(m)
	}
}
