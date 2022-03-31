package main

import "time"

/*
	1. 运行时会对字典并发操作做出检测。如果某个任务正在对字典进行写操作，那么其他任务就不能对该字典执行并发操作（读、写、删除），否则会导致进程奔溃。
*/

func main() {
	m := make(map[string]int)
	go func() {
		for {
			m["a"] += 1 // 写操作
			time.Sleep(time.Microsecond)
		}
	}()
	go func() {
		for {
			_ = m["b"] // 读操作
			time.Sleep(time.Microsecond)
		}
	}()
	select {} // 阻止进程退出
}
