/*
   @Author: StudentCWZ
   @Description:
   @File: main
   @Software: GoLand
   @Project: GoNotebook
   @Date: 2022/5/12 09:10
*/

package main

import (
	"fmt"
	"runtime"
	"time"
)

/*
	1. 如果在 main.main 里调用 Goexit，它会等待其他任务结束，然后让进程直接崩溃。
*/

func main() {
	for i := 0; i < 2; i++ {
		go func(x int) {
			for n := 0; n < 2; n++ {
				fmt.Printf("%c: %d\n", 'a'+x, n)
				time.Sleep(time.Millisecond)
			}
		}(i)
	}

	runtime.Goexit()
	println("main exit.") // 等待所有结果结束
}
