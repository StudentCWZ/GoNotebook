/*
   @Author: StudentCWZ
   @Description:
   @File: main
   @Software: GoLand
   @Project: GoNotebook
   @Date: 2022/5/16 11:22
*/

package main

import (
	"sync"
	"time"
)

/*
	1. 通知可以是群体性的，也未必就是通知结束，可以是任何需要表达的事件
*/

func main() {
	var wg sync.WaitGroup
	ready := make(chan struct{})

	for i := 0; i < 3; i++ {
		wg.Add(1)

		go func(id int) {
			defer wg.Done()

			println(id, ": ready.") // 运动员准备就绪
			<-ready                 // 等待发指令
			println(id, ": running ...")
		}(i)

		time.Sleep(time.Second)
		println("ready? Go!")

		close(ready)

		wg.Wait()

	}

}
