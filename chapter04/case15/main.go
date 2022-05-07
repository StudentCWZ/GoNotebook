/*
   @Author: StudentCWZ
   @Description:
   @File: main
   @Software: GoLand
   @Project: GoNotebook
   @Date: 2022/5/7 13:47
*/

package main

import "runtime/debug"

/*
	1. 调试阶段，可使用 runtime/debug.PrintStack() 函数输出完整调用堆栈信息。
	2. 建议：除非是不可恢复性、导致系统无法正常工作的错误，否则不建议使用 panic 。
*/

func test() {
	panic("i am dead")
}

func main() {
	defer func() {
		if err := recover(); err != nil {
			debug.PrintStack()
		}
	}()

	test()
}
