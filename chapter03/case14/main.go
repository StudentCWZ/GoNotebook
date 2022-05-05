/*
   @Author: StudentCWZ
   @Description:
   @File: main
   @Software: GoLand
   @Project: GoNotebook
   @Date: 2022/5/5 11:31
*/

package main

/*
	1. 仅有 for 一种循环语句，但常见方式都能支持。
	2. 初始化语句仅被执行一次。条件表达式中如有函数调用，需确认是否会重复执行。可能会被编译器优化掉，也可能是动态结果须每次执行确认。
*/

func count() int {
	print("count.")
	return 3
}

func forCaseOne() {
	for i := 0; i < 3; i++ {
		println(i) // 初始化表达式支持函数调用或定义局部变量
	}
}

func forCaseTwo() {
	x := 5
	for x < 10 { // 类似 while x < 10 {}
		x++
		println(x)
	}

	for {
		break // 类似 while true {}
	}
}

func forCaseThree() {
	for i, c := 0, count(); i < c; i++ { // 初始化语句的 count 函数仅执行一次
		println("a", i)
	}
	c := 0
	for c < count() { // 条件表达式中 count 重复执行
		println("b", c)
		c++
	}
	println("b", c)
}

func main() {
	forCaseOne()
	forCaseTwo()
	forCaseThree()

}
