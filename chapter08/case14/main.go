/*
   @Author: StudentCWZ
   @Description:
   @File: main
   @Software: GoLand
   @Project: GoNotebook
   @Date: 2022/5/12 10:08
*/

package main

/*
	1. 内置函数 cap 和 len 返回缓冲区大小和当前已缓冲数量；而对于同步通道则都返回 0，据此可判断通道是同步还是异步。
*/

func main() {
	a, b := make(chan int), make(chan int, 3)
	b <- 1
	b <- 2
	println("a:", len(a), cap(a))
	println("b:", len(b), cap(b))
}
