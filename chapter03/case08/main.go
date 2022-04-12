/*
   @Author: StudentCWZ
   @Description:
   @File: main
   @Software: GoLand
   @Project: GoNotebook
   @Date: 2022/4/12 10:06
*/

package main

/*
	1. 指针没有专门指向成员的 "->" 运算符，统一使用 "." 选择表达式。
*/

func main() {
	a := struct {
		x int
	}{}

	a.x = 100
	p := &a

	p.x += 100 // 相当于 p -> x += 100
	println(p.x)

}
