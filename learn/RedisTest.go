package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

// 测试连接redis
func main() {
	conn, _ := redis.Dial("tcp", "192.168.137.200:6379")
	defer conn.Close()
	// 如果不存在就设置
	_, _ = conn.Do("SETNX", "DABAO", "666")
	_, _ = conn.Do("HSET", "USER_ACCOUNT", "USER_NAME", "过江龙")
	// 设置过期时间
	_, _ = conn.Do("EXPIRE", "USER_ACCOUNT", 100)
	_, _ = conn.Do("EXPIRE", "DABAO", 100)
	// 取数据
	do, _ := conn.Do("HGET", "USER_ACCOUNT", "USER_NAME")
	uint8s := do.([]uint8)
	fmt.Println(string(uint8s))
}
