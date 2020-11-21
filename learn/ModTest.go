package main

import "core/utils"

// 引入第三方包或者自己写的包
// 第一步: 在项目根目录下进行命令执行-> go mod init xxx(随意定)
// 第二步: 创建代码目录
// 第三部: 在代码目录下编写自己的工具类
// 在对应的main中进行使用
func main() {
	result := utils.Add(1, 1)
	utils.Println(result)
	utils.Println(utils.Age)
}
