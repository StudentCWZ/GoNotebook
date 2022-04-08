/*
   @Author: StudentCWZ
   @Description:
   @File: main
   @Software: GoLand
   @Project: GoNotebook
   @Date: 2022/4/8 09:59
*/

package main

/*
	1. 可以为当前包内的任意类型定义方法
*/

type X int

func (x *X) inc() { //	名称前的参数称作 receiver，作用类似 python self。
	*x++
}

func main() {
	var x X

	x.inc()
	println(x)
}
