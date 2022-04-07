/*
   @Author: StudentCWZ
   @Description:
   @File: main
   @Software: GoLand
   @Project: GoNotebook
   @Date: 2022/4/7 11:08
*/

package main

import "fmt"

/*
	1. 定义函数类型，让相同签名的函数自动实现某个接口。
*/

type FuncString func() string

func (f FuncString) String() string {
	return f()
}

func main() {
	var t fmt.Stringer = FuncString(func() string {
		return "hello, world!"
	})

	fmt.Println(t)
}