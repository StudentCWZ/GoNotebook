/*
   @Author: StudentCWZ
   @Description:
   @File: main
   @Software: GoLand
   @Project: GoNotebook
   @Date: 2022/4/12 08:44
*/

package main

import "fmt"

/*
	1. Go 语言仅仅 25 个保留关键字，这是最常见的宣传语，虽然不是主流语言中最少的，但也确实体现了 Go 语法规则的简洁性。
	2. 保留关键字不能用作常量、变量、函数名，以及结构字段等标识符。
	3. 运算符和表达式用来串联数据和指令，算是最基础的算法。
	4. 没有乘幂和绝对值运算符，对应的是标准库 math 里的 Pow、Abs 函数实现。
	5. 一元运算符优先级最高，二元则分成五个级别，相同优先级的二元运算符，从左往右依次计算。
	6. 除位移操作外，操作数类型必须相同。如果其中一个是无显式类型声明的常量，那么该常量的操作数会自动转型。
*/

func main() {
	const v = 20 // 无显式类型声明的常量
	fmt.Printf("%T, %v\n", v, v)

	var a byte = 10
	b := v + a // v 自动转换为 byte/uint8 类型
	fmt.Printf("%T, %v\n", b, b)

	const c float32 = 1.2
	d := c + v // v 自动转换为 float32 类型
	fmt.Printf("%T, %v\n", d, d)
}
