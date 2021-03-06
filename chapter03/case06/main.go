/*
   @Author: StudentCWZ
   @Description:
   @File: main
   @Software: GoLand
   @Project: GoNotebook
   @Date: 2022/4/12 09:39
*/

package main

import "fmt"

/*
	1. 不能将内存地址与指针混为一谈
	2. 内存地址是内存中每个字节单元的唯一编号，而指针则是一个实体。指针会分配内存空间，相当于一个专门用来保存地址的整型变量。
       - 取址运算符 "&" 用于获取对象地址
	   - 指针运算符 "*" 用于间接引用目标对象。
	   - 二级指针 **T，如包含包名则写成 *package.T。
	3. 并非所有对象都能进行取地址操作，但变量总是能正确返回。
	4. 指针运算符为左值时，我们可更新目标对象状态；而为右值时则是为了获取目标状态。
*/

func main() {
	x := 10
	var p *int = &x // 获取地址，保存指针变量
	*p += 20        // 用指针间接引用，并更新对象

	println(p, *p) // 输出指针所存储的地址，以及目标地址

	m := map[string]int{"a": 1}
	// println(&m["a"])	// 报错
	fmt.Printf("%+v\n", m)
}
