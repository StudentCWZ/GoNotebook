/*
   @Author: StudentCWZ
   @Description:
   @File: main
   @Software: GoLand
   @Project: GoNotebook
   @Date: 2022/4/7 09:12
*/

package main

/*
	1. 如果接口没有任何方法声明，那么就是一个空接口(interface{})，它的用途类似面向对象里面的根类型 Object，可被赋值为任何类型的对象。
	2. 接口变量默认值是 nil。如果实现接口的类型支持，可做相等运算
*/

func main() {
	var t1, t2 interface{}
	println(t1 == nil, t2 == nil)

	t1, t2 = 100, 100
	println(t1 == t2)

	//t1, t2 = map[string]int{}, map[string]int{}	报错
	//println(t1 == t2)
}
