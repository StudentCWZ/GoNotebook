/*
   @Author: StudentCWZ
   @Description:
   @File: main
   @Software: GoLand
   @Project: GoNotebook
   @Date: 2022/4/6 10:03
*/

package main

/*
	1. 指针类型的 receiver 必须是合法指针(包括 nil)，或能获取实例地址。
	2. 如何选择方法的 receiver 类型
		- 要修改实例状态，用 *T
		- 无须修改状态的小对象或固定值，建议用 T。
		- 大对象建议用 *T，以减少复制成本。
		- 引用类型、字符串、函数等指针包装对象，直接用 T。
		- 若包含 Mutex 等同步字段，用 *T，避免因复制造成锁操作无效。
		- 其他无法确定的情况，都用 *T。
*/

type X struct{}

func (x *X) test() {
	println("hi!", x)
}

func main() {
	var a *X
	a.test() // 相当于 test(nil)
	//X{}.test()	// 报错：cannot take the address of X literal
}
