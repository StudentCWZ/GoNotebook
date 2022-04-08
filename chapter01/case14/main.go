/*
   @Author: StudentCWZ
   @Description:
   @File: main
   @Software: GoLand
   @Project: GoNotebook
   @Date: 2022/4/8 10:49
*/

package main

import (
	"fmt"
	"time"
)

/*
	1. 整个运行时完全并发化设计。凡你能看到的，几乎都在以 goroutine 方式运行。
	2. 这是一种比普通协程或线程更加高效的并发设计，能轻松创建和运行成千上万的并发任务。
*/

func task(id int) {
	for i := 0; i < 5; i++ {
		fmt.Printf("%d: %d\n", id, i)
		time.Sleep(time.Second)
	}
}

func main() {
	go task(1) // 创建 goroutine
	go task(2)

	time.Sleep(time.Second * 6)
}
