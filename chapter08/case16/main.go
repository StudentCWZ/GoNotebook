/*
   @Author: StudentCWZ
   @Description:
   @File: main
   @Software: GoLand
   @Project: GoNotebook
   @Date: 2022/5/16 11:15
*/

package main

/*
	1. 对于循环接收数据，range 模式更加简洁一些
*/

func main() {
	done := make(chan struct{})
	c := make(chan int)

	go func() {
		defer close(done)

		for x := range c { // 循环获取消息，直到通道被关闭
			println(x)
		}
	}()

	c <- 1
	c <- 2
	c <- 3
	close(c)

	<-done
}
