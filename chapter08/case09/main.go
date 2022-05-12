/*
   @Author: StudentCWZ
   @Description:
   @File: main
   @Software: GoLand
   @Project: GoNotebook
   @Date: 2022/5/12 08:50
*/

package main

import "runtime"

/*
	1. Goexit 立即终止当前任务，运行时确保所有已注册延迟调用被执行。该函数不会影响其他并发任务，不会引起 panic，自然也就无法捕获。
*/

func main() {
	exit := make(chan struct{})

	go func() {
		defer close(exit)  // 执行
		defer println("a") // 执行
		func() {
			defer func() {
				println("b", recover() == nil) // 执行，recover 返回 nil
			}()
			func() { // 在多层调用中执行 Goexit
				println("c")
				runtime.Goexit()   // 立即终止整个调用堆栈
				println("c done.") // 不会执行
			}()
			println("b done.") // 不会执行
		}()
		println("a done.") // 不会执行
	}()
	<-exit
	println("main exit.")
}
