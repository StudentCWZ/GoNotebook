/*
   @Author: StudentCWZ
   @Description:
   @File: main
   @Software: GoLand
   @Project: GoNotebook
   @Date: 2022/4/12 09:23
*/

package main

import "fmt"

/*
	1. 二进制位运算符比较特别的就是 "bit clear"，在其他语言里很少见到。
	2. 位清除和位亦或是不同的，它将左右操作数对应的二进制位都为 1 的重置为 0，以达到一次清楚多个标记位的目的。
*/

const (
	read byte = 1 << iota
	write
	exec
	freeze
)

func main() {
	a := read | write | freeze
	b := read | freeze | exec
	c := a &^ b

	fmt.Printf("%04b &^ %04b = %04b\n", a, b, c)
}
