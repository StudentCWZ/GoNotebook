/*
   @Author: StudentCWZ
   @Description:
   @File: main
   @Software: GoLand
   @Project: GoNotebook
   @Date: 2022/5/11 16:47
*/

package main

import (
	"fmt"
	"sync"
)

/*
	1. 与线程不同，goroutine 任务无法设置优先级，无法获取编号，没有局部存储(TLS)，设置连返回值都会被抛弃。但除优先级外，其他功能都很容易实现。
*/

func main() {
	var wg sync.WaitGroup
	var gs [5]struct { // 用于实现类似 TLS 功能
		id     int // 编号
		result int // 返回值
	}

	for i := 0; i < len(gs); i++ {
		wg.Add(1)

		go func(id int) { // 使用参数避免闭包延迟求值
			defer wg.Done()

			gs[id].id = id
			gs[id].result = (id + 1) * 100
		}(i)
	}

	wg.Wait()
	fmt.Printf("%+v\n", gs)
}
