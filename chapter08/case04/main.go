/*
   @Author: StudentCWZ
   @Description:
   @File: main
   @Software: GoLand
   @Project: GoNotebook
   @Date: 2022/5/11 15:55
*/

package main

import (
	"sync"
	"time"
)

/*
	1. 如要等待多个任务结束，推荐使用 sync.WaitGroup。通过设定计数器，让每个 goroutine 在退出前递减，直至归零时解除阻塞。
*/

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1) // 累加计数

		go func(id int) {
			defer wg.Done()

			time.Sleep(time.Second)
			println("goroutine", id, "done.")
		}(i)
	}
	println("main ...")
	wg.Wait() // 阻塞，知道计数归零
	println("main exit.")
}
