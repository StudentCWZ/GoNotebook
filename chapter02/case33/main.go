/*
   @Author: StudentCWZ
   @Description:
   @File: main.go
   @Software: GoLand
   @Project: GoNotebook
   @Date: 2022/4/10 12:09
*/

package main

import "fmt"

/*
	1. 与明确标识符的 bool、int、string 类型相比，数组、切片、字典、通道等类型与具体元素类型或长度等属性有关，故称作未命名类型。
	2. 但个案，可用 type 为其提供具体名称，将其改变为命名类型。
	3. 具有相同声明的未命名类型被视作同一类型：
       - 具有相同基类型的指针
       - 具有相同元素类型和长度的数组
       - 具有相同元素类型的切片
       - 具有相同键值类型的字典
       - 具有相同数据类型及操作方向的通道
       - 具有相同字段序列的结构体
       - 具有相同签名的函数
       - 具有相同的方法集的接口
	4. 容易被忽视得失 struct tag，它也属于类型组成部分，而不仅仅是元数据描述。
*/

func structTest() {
	var a struct { // 匿名结构体
		x int    `x`
		s string `s`
	}

	var b struct {
		x int
		s string
	}
	//b = a // 报错
	fmt.Printf("%+v\n", a)
	fmt.Printf("%+v\n", b)
}

func funcTest() {
	var a func(int, string) int
	var b func(string, int) int
	//b = a // 报错
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
}

func main() {
	structTest()
	funcTest()
}
