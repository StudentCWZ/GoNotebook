/*
   @Author: StudentCWZ
   @Description:
   @File: main.go
   @Software: GoLand
   @Project: GoNotebook
   @Date: 2022/4/9 21:36
*/

package main

import "fmt"

/*
	1. Go 并没有明确意义上的 enum 的定义，不过可借助 iota 标识符实现一组自增常量来实现枚举类型。
*/

func main() {
	const (
		x = iota // 0
		y        // 1
		z        // 2
	)
	const (
		_  = iota             // 0
		KB = 1 << (10 * iota) // 1 << (10 * 1)
		MB                    //	1 << (10 * 2)
		GB                    // 1 << (10 * 3)
	)

	fmt.Println(x, y, z)
	fmt.Println(KB, MB, GB)
}
