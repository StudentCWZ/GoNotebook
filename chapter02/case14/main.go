/*
   @Author: StudentCWZ
   @Description:
   @File: main.go
   @Software: GoLand
   @Project: GoNotebook
   @Date: 2022/4/9 21:29
*/

package main

import "fmt"

/*
	1.在常量组中如不制定类型和初始化值，则与上一行非空常量右值(表达式文本)相同。
*/

func main() {
	const (
		x uint16 = 120
		y        // 与上一行 x 类型、右值相同
		s = "abc"
		z // 与 s 类型、右值相同
	)

	fmt.Printf("%T, %v\n", y, y)
	fmt.Printf("%T, %v\n", z, z)
}
