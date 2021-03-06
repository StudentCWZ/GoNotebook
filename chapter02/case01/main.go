/*
   @Author: StudentCWZ
   @Description:
   @File: main
   @Software: GoLand
   @Project: GoNotebook
   @Date: 2022/4/9 13:59
*/

package main

import "fmt"

/*
	1. 在数学概念中，变量表示没有固定值且可改变的数。但从计算机系统实现的角度来看，变量是一段或多段用来存储数据的内存。
	2. 作为静态类型语言，Go 变量总是有固定的数据类型，类型决定了变量内存的长度和存储格式。
	3. 我们只能修改变量值，无法改变类型。
	4. 通过类型转换或指针操作，我们可用不同方式修改变量的值，这并不意味着改变了变量类型。
	5. 因为内存分配发生在运行期，所以在编码阶段我们用一个易于阅读的名字来表示这段内存。
	6. 实际上，编译后的机器码从不使用变量名，而是直接通过内存地址来访问目标数据。
	7. 保存在符号表中的变量名等信息可被删除，或用于输出更详细的错误信息。
	8. 关键字 var 用于定义变量，和 C 不同，类型被放在变量名后。
	9. 另外，运行时内存分配操作会确保变量自动初始化为二进制零值，避免出现不可预测行为。
	10. 如果显式提供初始化值，可省略变量类型，由编译器进行推断。
*/

func main() {
	var x int     // 自动初始化为 0
	var y = false // 自动推断为 bool 类型

	// 可一次性定义多个变量，包含不同初始值定义不同类型
	var w, z int          // 相同类型的多个变量
	var a, s = 100, "abc" // 不同类型初始化值

	// 依照惯例，建议以组方式整理多行变量定义
	var (
		b, c int
		u, v = 100, "abc"
	)
	fmt.Println(x, y, w, z, a, s, b, c, u, v)
}