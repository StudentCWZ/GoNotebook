/*
   @Author: StudentCWZ
   @Description:
   @File: main.go
   @Software: GoLand
   @Project: GoNotebook
   @Date: 2022/4/10 12:03
*/

package main

import "fmt"

/*
	1. 即便指定了基础类型，也只表明它们有相同底层数据结构，两者间不存在任何关系，属完全不同的两种类型。
	2. 除操作符外，自定义类型不会继承基础类型的其他信息(包括方法)。不能视作别名，不能隐式转换，不能直接用于比较表达式。
*/

func main() {
	type data int
	var d data = 10
	//var x int = d	// 报错
	//println(d == x) // 报错
	fmt.Println(d)
}
