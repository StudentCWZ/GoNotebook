/*
   @Author: StudentCWZ
   @Description:
   @File: main
   @Software: GoLand
   @Project: GoNotebook
   @Date: 2022/5/9 16:04
*/

package main

import "time"

/*
	1. 关键字 go 并非执行并发操作，而是创建一个并发任务单元。新建任务被放置在系统队列中，等待调度器安排合适系统线程去获取执行权。当前流程不会
       阻塞，不会等待该任务启动，且运行时也不保证并发任务的执行次序。
	2. 每个任务单元除保存函数指针、调用参数外，还会分配执行所需的栈内存空间。相比系统默认 MB 级别的线程栈，goroutine 自定义栈仅须 2 KB，所
       以才能创建成千上万的并发任务，自定义栈采取按需分配策略，在需要时进行扩容，最大能到 GB 规模。
	3. 与 defer 一样，goroutine 也会因 "延迟执行" 而立刻计算并复制执行参数。
*/

var c int

func counter() int {
	c++
	return c
}

func main() {
	a := 100

	go func(x, y int) {
		time.Sleep(time.Second) // 让 goroutine 在 main 逻辑之后执行
		println("go:", x, y)
	}(a, counter()) // 立即计算并复制参数

	a += 100
	println("main:", a, counter())
	time.Sleep(time.Second * 3) // 等待 goroutine 结束
}
