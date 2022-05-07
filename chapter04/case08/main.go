/*
   @Author: StudentCWZ
   @Description:
   @File: main
   @Software: GoLand
   @Project: GoNotebook
   @Date: 2022/5/7 10:16
*/

package main

import (
	"fmt"
	"log"
	"os"
)

/*
	1. 语句 defer 向当前函数注册稍后执行的函数调用。这些调用被称作延迟调用，因为它们直到当前函数执行结束前才执行，常用于资源释放、解除锁定，
       以及错误处理等操作。
	2. 注意，延迟调用注册的是调用，必须提供执行所需参数(哪怕为空)。参数值在注册时被复制并缓存起来。如对状态敏感，可改用指针或闭包。
	3. 多个延迟注册按 FILO 次序执行。
	4. 编译器通过插入额外指令来实现延迟调用执行，而 return 和 panic 语句都会终止当前函数流程，引发延迟调用。另外，return 语句不是 ret 汇
       编指令，它会先更新返回值。
	5. 千万记住，延迟调用在函数结束时才被执行。不合理的使用方式会浪费更多资源，甚至造成逻辑错误。
*/

func deferCaseOne() {
	f, err := os.Open("../case07/main.go")
	if err != nil {
		log.Fatal(err)
	}

	defer func(f *os.File) { // 仅注册，直到 deferCaseOne 退出前才执行
		err := f.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(f)
}

func deferCaseTwo() {
	x, y := 1, 2

	defer func(a int) { // y 为 闭包引用
		println("defer x, y = ", a, y) // 注册时复制调用参数
	}(x)

	x += 100 // 对 x 的修改不会影响延迟调用
	y += 200
	println(x, y)
}

func deferCaseThree() {
	defer println("a")
	defer println("b")
}

func deferCaseFour() (z int) {
	defer func() {
		println("defer:", z)
		z += 100
	}()
	return 100
}

func deferCaseFive() {
	do := func(n int) {
		path := fmt.Sprintf("../case0%d/main.go", n)
		f, err := os.Open(path)
		if err != nil {
			log.Println(err)
		}
		defer func(f *os.File) {
			err := f.Close()
			if err != nil {
				log.Println(err)
			}
		}(f)
	}
	for i := 1; i < 5; i++ {
		do(i)
	}
}

func main() {
	deferCaseOne()
	deferCaseTwo()
	deferCaseThree()
	println("test:", deferCaseFour())
	deferCaseFive()
}
