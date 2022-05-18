/*
   @Author: StudentCWZ
   @Description:
   @File: main
   @Software: GoLand
   @Project: GoNotebook
   @Date: 2022/5/18 10:44
*/

package main

import "sync"

/*
	1. 通道默认是双向的，并不区分发送和接收端。但某些时候，我们可限制收发操作的方向来获取更严谨的操作逻辑。
	2. 尽管可用 make 创建单向通道，但那没有任何意义。通常使用类型转换来获取单项通道，并分别赋予操作双方。
*/

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	c := make(chan int)
	var send chan<- int = c
	var recv <-chan int = c

	go func() {
		defer wg.Done()

		for x := range recv {
			println(x)
		}
	}()

	go func() {
		defer wg.Done()
		defer close(c)

		for i := 0; i < 3; i++ {
			send <- i
		}
	}()

	wg.Wait()
}
