/*
   @Author: StudentCWZ
   @Description:
   @File: main.go
   @Software: GoLand
   @Project: GoNotebook
   @Date: 2022/4/10 00:00
*/

package main

import "fmt"

/*
	1. 拥有相同底层结构的并不一定就属于别名。就算在 64 位平台上 int 和 int64 结构完全一致，也分属不同类型，须显式转换。
*/

func add(x, y int) int {
	return x + y
}

func main() {
	var x int = 100
	//var y int64 = x // 报错
	var y int64 = 20
	z := int(y)
	fmt.Println(add(x, z))
}
