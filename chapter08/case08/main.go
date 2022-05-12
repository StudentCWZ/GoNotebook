/*
   @Author: StudentCWZ
   @Description:
   @File: main
   @Software: GoLand
   @Project: GoNotebook
   @Date: 2022/5/12 08:35
*/

package main

import "runtime"

/*
	1. 暂停，释放线程去执行其他任务。当前任务被放回队列，等待下次调度时恢复执行。
*/

func main() {
	runtime.GOMAXPROCS(1)
	exit := make(chan struct{})

	go func() { // 任务 a
		defer close(exit)

		go func() { // 任务 b 放在此处，是为了确保 a 优先执行
			println("b")
		}()

		for i := 0; i < 4; i++ {
			println("a:", i)
			if i == 1 { // 让出当前线程，调度执行 b
				runtime.Gosched()
			}
		}
	}()
	<-exit
}
