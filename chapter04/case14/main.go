/*
   @Author: StudentCWZ
   @Description:
   @File: main
   @Software: GoLand
   @Project: GoNotebook
   @Date: 2022/5/7 13:42
*/

package main

/*
	1. 考虑到 recover 特性，如果要保护代码片段，那么只能将其重构为函数调用。
*/

func test(x, y int) {
	z := 0

	func() {
		defer func() {
			if recover() != nil {
				z = 0
			}
		}()
		z = x / y
	}()
	println("x / y =", z)
}

func main() {
	test(5, 0)
}
