/*
   @Author: StudentCWZ
   @Description:
   @File: main
   @Software: GoLand
   @Project: GoNotebook
   @Date: 2022/5/11 16:22
*/

package main

import (
	"math"
	"runtime"
	"sync"
)

/*
	1. 运行时可能会创建很多线程，但任何时候仅有限的几个线程参与并发任务执行。该数量默认与处理器核数相等，可用 runtime.GOMAXPROCS 函数修改。
*/

// 测试目标函数
func count() {
	x := 0
	for i := 0; i < math.MaxUint32; i++ {
		x += i
	}
	println(x)
}

// 循环执行
func test(n int) {
	for i := 0; i < n; i++ {
		count()
	}
}

// 并发执行
func test2(n int) {
	var wg sync.WaitGroup
	wg.Add(n)

	for i := 0; i < n; i++ {
		go func() {
			count()
			wg.Done()
		}()
	}

	wg.Wait()
}

func main() {
	n := runtime.GOMAXPROCS(0)
	test(n)
	//test2(n)
}
