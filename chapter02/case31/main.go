/*
   @Author: StudentCWZ
   @Description:
   @File: main.go
   @Software: GoLand
   @Project: GoNotebook
   @Date: 2022/4/10 11:56
*/

package main

import "fmt"

/*
	1. 和 var、const 类似，多个 type 定义可合并成组，可在函数或代码快内定义局部类型。
*/

func main() {
	type (
		user struct {
			name string
			age  uint8
		}
		event func(string) bool // 函数类型
	)

	u := user{"Tom", 20}
	fmt.Println(u)

	var f event = func(s string) bool {
		println(s)
		return s != ""
	}

	f("abc")
}