/*
   @Author: StudentCWZ
   @Description:
   @File: main
   @Software: GoLand
   @Project: GoNotebook
   @Date: 2022/4/6 11:06
*/

package main

import "sync"

/*
	1. 可以像访问匿名字段成员那样调用其方法，由编译器负责查找
*/

type data struct {
	sync.Mutex // guards
	buf        [1024]byte
}

func main() {
	d := data{}
	d.Lock() // 编译会处理为 sync.(*Mutex).Lock() 调用
	defer d.Unlock()
}
