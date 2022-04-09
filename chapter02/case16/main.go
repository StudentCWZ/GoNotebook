/*
   @Author: StudentCWZ
   @Description:
   @File: main.go
   @Software: GoLand
   @Project: GoNotebook
   @Date: 2022/4/9 21:43
*/

package main

import "fmt"

/*
	1. 自增作用范围为常量组。可在多常量定义中使用多个 iota，它们各自单独记数，只须确保组中每行常量的列数量相同即可。
*/

func main() {
	const (
		_, _ = iota, iota * 10 // 0, 0 * 10
		a, b                   // 1, 1 * 10
		c, d                   // 2, 2 * 10
	)
	fmt.Println(a, b, c, d)
}
