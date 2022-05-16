/*
   @Author: StudentCWZ
   @Description:
   @File: main
   @Software: GoLand
   @Project: GoNotebook
   @Date: 2022/5/16 11:45
*/

package main

/*
	1. 对于 closed 或 nil 通道，发送和接收操作都有相应规则：
	   - 向已关闭通道发送数据，引发 panic
       - 已关闭接收数据，返回已缓冲数据或零值。
       - 无论收发，nil 通道都会阻塞。
	2. 重复关闭，或关闭 nil 通道都会引发 panic 错误
*/

func main() {
	c := make(chan int, 3)

	c <- 10
	c <- 20
	c <- 30
	close(c)

	for i := 0; i < cap(c)+1; i++ {
		x, ok := <-c
		println(i, ":", ok, x)
	}
}
