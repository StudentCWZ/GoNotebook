/*
   @Author: StudentCWZ
   @Description:
   @File: main.go
   @Software: GoLand
   @Project: GoNotebook
   @Date: 2022/4/9 21:49
*/

package main

import "fmt"

/*
	1. 如中断 iota 自增，则必须显式恢复。且后续自增按行序递增，而非 C enum 按上一取值递增。
*/

func main() {
	const (
		a = iota // 0
		b        // 1
		c = 100  //	100
		d        // 100(与上一行常量右值表达式相同)
		e = iota // 4(恢复 iota 自增，计数包含 c、d)
		f        // 5
	)
	fmt.Println(a, d, c, d, e, f)
}
