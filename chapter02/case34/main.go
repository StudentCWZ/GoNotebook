/*
   @Author: StudentCWZ
   @Description:
   @File: main.go
   @Software: GoLand
   @Project: GoNotebook
   @Date: 2022/4/10 12:30
*/

package main

import "fmt"

/*
	1. 未命名类型转换规则
       - 所属类型相同
       - 基础类型相同，且其中一个是未命名类型
       - 数据类型相同，将双向通道赋值给单向通道，且其中一个为未命名类型
       - 将默认值 nil 赋值给切片、字典、通道、指针、函数或接口
       - 对象实现了目标接口
*/

func main() {
	type data [2]int
	var d data = [2]int{1, 2} // 基础类型相同，右值为未命名类型
	fmt.Println(d)

	a := make(chan int, 2) // 双向通道转换为单向通道，其中 b 为未命名类型
	var b chan<- int = a
	b <- 2
	fmt.Println(b)
}
