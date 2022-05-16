/*
   @Author: StudentCWZ
   @Description:
   @File: main
   @Software: GoLand
   @Project: GoNotebook
   @Date: 2022/5/16 11:01
*/

package main

/*
	1. 除使用简单的发送和接收操作外，还可用 ok-idom 或 range 模式处理数据。
*/

func main() {
	done := make(chan struct{})
	c := make(chan int)

	go func() {
		defer close(done) // 确保发出结束通知

		for {
			x, ok := <-c
			if !ok { // 据此判断通道是否被关闭
				return
			}

			println(x)
		}
	}()

	c <- 1
	c <- 2
	c <- 3
	close(c)
	<-done
}
