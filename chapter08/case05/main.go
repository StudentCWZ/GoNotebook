/*
   @Author: StudentCWZ
   @Description:
   @File: main
   @Software: GoLand
   @Project: GoNotebook
   @Date: 2022/5/11 16:09
*/

package main

import (
	"sync"
	"time"
)

/*
	1. 尽管 WaitGroup.Add 实现了原子操作，但建议在 goroutine 外累加计数器，以免 Add 尚未执行，Wait 已经退出。
	2. 可在多处使用 Wait 阻塞，它们都能接收到通知。
*/

func test() {
	var wg sync.WaitGroup

	go func() {
		wg.Add(1)
		println("hi!")
		println("--------------------")
	}()

	wg.Wait()
	println("exit.")
}

func main() {
	test()
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		wg.Wait() // 等待归零，解除阻塞
		println("wait exit.")
	}()

	go func() {
		time.Sleep(time.Second)
		println("done.")
		wg.Done() // 递减计数
	}()

	wg.Wait() // 等待归零，解除阻塞
	println("main exit.")
}
