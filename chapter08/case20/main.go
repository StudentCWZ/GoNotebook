/*
   @Author: StudentCWZ
   @Description:
   @File: main
   @Software: GoLand
   @Project: GoNotebook
   @Date: 2022/5/18 10:57
*/

package main

/*
	1. 不能在单向通道上做逆向操作
*/

func main() {
	c := make(chan int, 2)

	var send chan<- int = c
	var recv <-chan int = c

	//<-send // 报错
	//recv<- 1

	send <- 1
	<-recv
	println(send)
	println(recv)
}
