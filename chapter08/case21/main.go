/*
   @Author: StudentCWZ
   @Description:
   @File: main
   @Software: GoLand
   @Project: GoNotebook
   @Date: 2022/5/23 18:52
*/

package main

import "sync"

/*
	1. 如果要同时处理多个通道，可选用 select 语句。它会随机选择一个可用通道做收发操作。
*/

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	a, b := make(chan int), make(chan int)

	go func() { // 接收端
		defer wg.Done()

		for {
			var (
				name string
				x    int
				ok   bool
			)
			select {
			case x, ok = <-a:
				name = "a"
			case x, ok = <-b:
				name = "b"
			}
			if !ok { // 如果任一通道关闭，则终止接收
				return
			}
			println(name, x) // 输出接收的数据信息
		}
	}()

	go func() { // 发送端
		defer wg.Done()
		defer close(a)
		defer close(b)

		for i := 0; i < 10; i++ {
			select {
			case a <- i:
			case b <- i * 10:
			}
		}
	}()

	wg.Wait()

}
