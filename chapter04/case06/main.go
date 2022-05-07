/*
   @Author: StudentCWZ
   @Description:
   @File: main
   @Software: GoLand
   @Project: GoNotebook
   @Date: 2022/5/7 08:54
*/

package main

/*
	1. 匿名函数是指没有定义名字符号的函数。
	2. 除了没有名字外，匿名函数和普通函数完全相同。最大区别是，我们可在函数内部定义匿名函数，形成类似嵌套效果。
	3. 匿名函数可直接调用，保存到变量，作为参数或返回值。
	4. 将匿名函数赋值给变量，与为普通函数提供名字标识符有着根本区别。当然，编译器会为匿名函数生成一个 "随机" 符号名。
	5. 普通函数和匿名函数都可作为结构体字段，或经通道传递。
	6. 不曾使用的匿名函数会被编译器当作错误。
	7. 除闭包因素外，匿名函数也是一种常见重构手段。可将大函数分解成多个相对独立的匿名函数块，然后用相对简洁的调用完成逻辑流程，以实现框架和细节
       分离。
    8. 相比语句块，匿名函数的作用被隔离(不使用闭包)，不会引发外部污染，更加灵活。没有定义顺序限制，必要时可抽离，便于实现干净、清晰的代码层次。
*/

func caseOne() {
	func(s string) {
		println(s)
	}("hello, world!")
}

func caseTwo() {
	add := func(x, y int) int {
		return x + y
	}
	println(add(1, 2))
}

func caseThree() func(int, int) int {
	return func(x int, y int) int {
		return x + y
	}
}

func testStruct() {
	type calc struct {
		mul func(x, y int) int
	}

	z := calc{
		mul: func(x, y int) int {
			return x * y
		},
	}
	println(z.mul(2, 3))
}

func testChannel() {
	c := make(chan func(int, int) int, 2)
	c <- func(x int, y int) int {
		return x + y
	}
	println((<-c)(1, 2))

}

func main() {
	caseOne()
	caseTwo()
	add := caseThree()
	println(add(1, 2))
	testStruct()
	testChannel()

}
