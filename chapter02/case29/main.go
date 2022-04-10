/*
   @Author: StudentCWZ
   @Description:
   @File: main.go
   @Software: GoLand
   @Project: GoNotebook
   @Date: 2022/4/10 11:27
*/

package main

import "fmt"

/*
	1. 如果转换的目标是指针、单向通道或没有返回值函数类型，那么必须使用括号，以避免造成语法分解错误。
	2。 正确的做法是用括号，让编译器将 *int 解析为指针类型
*/

func main() {
	x := 100
	//p := *int(&x) // 报错
	p := (*int)(&x)
	fmt.Println(p)
}
