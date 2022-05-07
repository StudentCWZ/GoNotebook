/*
   @Author: StudentCWZ
   @Description:
   @File: main
   @Software: GoLand
   @Project: GoNotebook
   @Date: 2022/5/7 09:37
*/

package main

/*
	1. 闭包是在其词法上下文中引用了自由变量的函数，或者说函数和其引用的环境的组合体。
	2. 闭包通过指针引用环境变量，那么可能会导致其生命周期延长，甚至被分配到堆内存。另外，还有所谓 "延迟求值" 的特性。
	3. 多个匿名函数引用同一环境变量，也会让事情变得更加复杂。任何的修改行为都会影响其他函数取之，在并发模式下可能需要做同步处理。
	4. 闭包让我们不用传递参数就可读取或修改环境状态，当然也要为此付出额外代价。对性能要求较高的场合，须慎重使用。
*/

func caseOne(x int) func() {
	return func() {
		println(x)
	}
}

func caseTwo(x int) func() {
	println(&x)
	return func() {
		println(&x, x)
	}
}

func caseThree() []func() {
	var s []func()

	for i := 0; i < 2; i++ {
		s = append(s, func() { // 将多个匿名函数添加到列表
			println(&i, i)
		})
	}
	return s // 返回匿名函数列表
}

func caseFour() []func() {
	var s []func()

	for i := 0; i < 2; i++ {
		x := i // x 每次循环都重新定义
		s = append(s, func() {
			println(&x, x)
		})
	}
	return s
}

func caseFive(x int) (func(), func()) { // 返回两个匿名函数
	return func() { // 修改环境变量
			println(x)
			x += 10
		}, func() { // 显示环境变量
			println(x)
		}
}

func main() {
	f := caseOne(123)
	f()
	f1 := caseTwo(0x100)
	f1()
	for _, f2 := range caseThree() { // 迭代执行所有匿名函数
		f2()
	}
	for _, f3 := range caseFour() {
		f3()
	}
	a, b := caseFive(100)
	a()
	b()
}
