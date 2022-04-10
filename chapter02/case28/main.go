/*
   @Author: StudentCWZ
   @Description:
   @File: main.go
   @Software: GoLand
   @Project: GoNotebook
   @Date: 2022/4/10 11:16
*/

package main

import "fmt"

/*
	1. 隐式转换造成的问题远大于它带来的好处。
	2. 除常量、别名类型以及未命名类型外，Go 强制要求使用显式类型转换。
	3. 加上不支持操作符重载，所以我们总是能确定语句及表达式的明确含义。
	4. 同样不能将非 bool 类型结果当作 true/false 使用。
*/

func testOne() {
	a := 10
	b := byte(a)
	//c := a + int(b)	// 混合类型表达式必须确保类型一致
	fmt.Println(a, b)
}

func testTwo() {
	x := 100
	//var b bool = x // 报错
	//if x {} // 报错
	fmt.Println(x)
}

func main() {
	testOne()
	testTwo()
}
