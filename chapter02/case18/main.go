/*
   @Author: StudentCWZ
   @Description:
   @File: main.go
   @Software: GoLand
   @Project: GoNotebook
   @Date: 2022/4/9 21:56
*/

package main

import "fmt"

/*
	1. 自增默认数据类型为 int，可显式指定类型。
*/

func main() {
	const (
		a         = iota // int
		b float32 = iota // float32
		c         = iota // int(如果不显式指定 iota, 则与 b 数据类型相同)
	)
	fmt.Println(a, b, c)
	fmt.Printf("a 的类型: %T, a 的值: %v\n", a, a)
	fmt.Printf("b 的类型: %T, b 的值: %v\n", b, b)
	fmt.Printf("c 的类型: %T, c 的值: %v\n", c, c)

}
