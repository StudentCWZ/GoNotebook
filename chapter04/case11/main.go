/*
   @Author: StudentCWZ
   @Description:
   @File: main
   @Software: GoLand
   @Project: GoNotebook
   @Date: 2022/5/7 11:22
*/

package main

import "log"

/*
	1. 与 error 相比， panic/recover 在使用方法上更接近 try/catch  结构化异常。
	2. 比较有趣的是，它们内置函数而非语句。panic 会立即中断当前函数流程，执行延时调用。而在延时调用函数中， recover 可捕获并返回 panic 提
       交的错误对象。
	3. 因为 panic 参数是空接口类型，因此可使用任何对象作为错误状态。而 recover 返回结果同样要做转型才能获得具体信息。
	4. 无论是否执行 recover，所有延迟调用都会执行，但中断性错误会沿调用栈向外传递，要么被外层捕获，要么导致进程崩溃。
*/

func panicCaseOne() {
	defer func() {
		if err := recover(); err != nil { // 捕获错误
			log.Fatalln(err)
		}
	}()
	panic("i am dead") // 引发错误
	println("exit.")   // 永远不会执行
}

func panicCaseTwo() {
	defer println("test.1")
	defer println("test.2")

	panic("i am dead")
}

func main() {
	//panicCaseOne()
	defer func() {
		log.Println(recover())
	}()

	panicCaseTwo()
}
