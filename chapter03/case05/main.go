/*
   @Author: StudentCWZ
   @Description:
   @File: main
   @Software: GoLand
   @Project: GoNotebook
   @Date: 2022/4/12 09:31
*/

package main

/*
	1. 自增、自减不再是运算符。只能作为独立语句，不能用于表达式。
	2. 表达式通常是求值代码，可作为右值或参数使用。而语句完成一个行为，比如 if、for 代码块。表达式可作为语句用，但语句却不能当作表达式。
*/

func main() {
	a := 1
	// ++a	// 语法错误
	// if (a++ > 1) {}	语法错误
	p := &a
	*p++ // 相当于 (*p)++
	println(a)
}
