/*
   @Author: StudentCWZ
   @Description:
   @File: main
   @Software: GoLand
   @Project: GoNotebook
   @Date: 2022/5/12 09:47
*/

package main

/*
	1. 同步模式必须有配对操作的 goroutine 出现，否则会一直堵塞。而异步模式在缓冲区未满或数据未读完前不会阻塞。
	2. 多数时候，异步通道有助于提升性能，减少排队阻塞。
*/

func main() {
	c := make(chan int, 3) // 创建带 3 个缓冲槽的异步通道
	c <- 1                 // // 缓冲区未满，不会阻塞
	c <- 2
	println(<-c)
	println(<-c) // 缓冲区尚有数据，不会阻塞
}
