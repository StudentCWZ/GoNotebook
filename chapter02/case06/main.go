/*
   @Author: StudentCWZ
   @Description:
   @File: main
   @Software: GoLand
   @Project: GoNotebook
   @Date: 2022/4/9 14:40
*/

package main

import (
	"fmt"
	"os"
)

/*
	1. 在处理函数错误返回值，退化赋值允许我们重复使用 err 变量，这是相当有益的。
*/

func main() {
	f, err := os.Open("/dev/random")
	if err != nil {
		panic(err)
	}
	buf := make([]byte, 1024)
	n, err := f.Read(buf) // err 退化赋值，n 新定义
	if err != nil {
		panic(err)
	}
	fmt.Println(n)
}
