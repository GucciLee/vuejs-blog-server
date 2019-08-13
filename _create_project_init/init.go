package main

import (
	"fmt"
	"vuejs-blog-server/utils/learnku_filepath"
)

// 创建环境配置文件
func createConfEnv()  {
	var env string = "../conf/env.conf"
	var envExample string = "../conf/env.example.conf"
	// 如果 evn.conf 文件不存在，就创建
	if !learnku_filepath.Exists(env) {
		printFmt("创建配置文件：", 1)
		_, err := learnku_filepath.CopyFile(envExample, env)
		if err != nil {
			fmt.Println(err)
		} else {
			printFmt("1）文件 " + env + " 创建成功，请进行配置 ？", 2)
		}
	}
}

// 格式化打印
func printFmt(v string, len int)  {
	var pre string
	for i:=0; i<len ;i++  {
		pre += "|--"
	}
	fmt.Println(pre + v)
}

func main() {
	fmt.Println("=========== 项目初始化开始。。。")
	createConfEnv()
	fmt.Println("=========== 项目初始化完成。。。")
}
