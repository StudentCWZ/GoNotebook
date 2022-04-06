/*
   @Author: StudentCWZ
   @Description:
   @File: main
   @Software: GoLand
   @Project: GoNotebook
   @Date: 2022/4/6 09:18
*/

package main

/*
	1. 方法同样不支持重载(overload)。
	2. receiver 参数名没有限制，按惯例会选用简短有意义的名称(不推荐使用 this、self)。
	3. 如方法内部并不引用实例，可省略参数名，仅保留类型。
*/

type N int

func (N) test() {
	println("hi!")
}

func main() {
	var n N
	n.test()
}
