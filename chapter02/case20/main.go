/*
   @Author: StudentCWZ
   @Description:
   @File: main.go
   @Software: GoLand
   @Project: GoNotebook
   @Date: 2022/4/9 22:11
*/

package main

/*
	1. 不同于变量在运行期分配存储内存(非优化状态)，常量通常会被编译器在预处理阶段直接展开，作为指定数据使用。
	2. 数字常量不会分配存储空间，无需像变量那样通过内存寻址来取值，因此无法获取地址。
*/

var x = 0x100

const y = 0x200

const a = 100    // 无类型声明的常量
const b byte = a // 直接展开 a，相当于 const b byte = 100

const u int = 100 // 显式指定常量类型，编译器会做强类型检查
//const v byte = u // 错误

func main() {
	println(&x, x)
	//println(&y, y) // 报错：cannot take the address of y
	println(y)
}
