/*
   @Author: StudentCWZ
   @Description:
   @File: main.go
   @Software: GoLand
   @Project: GoNotebook
   @Date: 2022/4/10 11:45
*/

package main

import "fmt"

/*
	1. 使用关键字 type 定义用户自定义类型，包括基于现有基础类型创建，或者是结构体、函数类型等。
*/

type flags byte

const (
	read flags = 1 << iota
	write
	exec
)

func main() {
	f := read | exec
	fmt.Printf("%b\n", read)
	fmt.Printf("%b\n", write)
	fmt.Printf("%b\n", exec)
	fmt.Printf("%b\n", f) // 输出二进制标记位
}
