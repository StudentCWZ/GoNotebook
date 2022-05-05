/*
   @Author: StudentCWZ
   @Description:
   @File: main
   @Software: GoLand
   @Project: GoNotebook
   @Date: 2022/5/5 09:01
*/

package main

import (
	"errors"
	"log"
)

/*
	1. Go 精简了流控制语句，虽然某些时候不够便捷，但够用。
	2. 条件表达式必须是布尔类型，可省略括号，且左花括号不能另起一行。
	3. 比较特别的是对初始化语句的支持，可定义你局部变量或执行初始化函数
	4. 局部变量的有效范围包含整个 if/else 块
	5. 对于编程初学者，可能会因条件匹配顺序不当而写出代码(dead code)。
	6. 死代码是指永远不会被执行的代码，可食用专门的工具，或用代码覆盖率测试进行检查。
	7. 某些比较智能的编译器也可主动清除死代码。
	8. 尽可能减少代码块嵌套，让正常逻辑处于相同层次。
*/

func xInit(x int) {
	x = 5
}

func check(x int) error {
	if x <= 0 {
		return errors.New("x <= 0")
	}
	return nil
}

func ifCaseOne() {
	x := 3
	if x > 5 {
		println("a")
	} else if x < 5 && x > 0 {
		println("b")
	} else {
		println("z")
	}
}

func ifCaseTwo() {
	x := 10

	if xInit(x); x == 0 { // 优先执行 xInit 函数
		println("a")
	} else {
		println("b")
	}

	if a, b := x+1, x+10; a < b { // 定义一个或多个局部变量(也可以是函数返回值)
		println("a")
	} else {
		println("b")
	}
}

func ifCaseThree() {
	x := 8

	if x > 5 { // 优先判断，条件表达式结果为 true
		println("a")
	} else if x > 7 { // dead code
		println("b")
	}
}

func ifCaseFour() {
	x := 10

	if err := check(x); err == nil {
		x++
		println(x)
	} else {
		log.Fatalln(err)
	}
}

func ifCaseFive() {
	x := 10

	if err := check(x); err != nil {
		log.Fatalln(err)
	}
	x++
	println(x)
}

func main() {
	ifCaseOne()
	ifCaseTwo()
	ifCaseThree()
	ifCaseFour()
	ifCaseFive()
}
