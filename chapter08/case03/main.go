/*
   @Author: StudentCWZ
   @Description:
   @File: main
   @Software: GoLand
   @Project: GoNotebook
   @Date: 2022/5/9 16:26
*/

package main

import "time"

/*
	1. 进程退出时，不会等待并发任务结束，可用通道(channel)阻塞，然后发出退出信号。
*/

func main() {
	exit := make(chan struct{}) // 创建通道。因为仅是通知，数据并没有实际意义。

	go func() {
		time.Sleep(time.Second)
		println("goroutine done")

		close(exit) // 关闭通道，发出信号
	}()

	println("main ...")
	<-exit // 如通道关闭，立刻解除阻塞
	println("main exit ...")
}
