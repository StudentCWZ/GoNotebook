/*
   @Author: StudentCWZ
   @Description:
   @File: main
   @Software: GoLand
   @Project: GoNotebook
   @Date: 2022/4/8 09:16
*/

package main

/*
	1. Go 仅有三种流控制语句，与大多数语言相比，都可称得上简单。
	2. 在迭代遍历时，for ... range 除元素外，还可以返回索引。
*/

// IfExpression (if 控制)
func IfExpression() {
	x := 100

	if x > 0 {
		println("x")
	} else if x < 0 {
		println("-x")
	} else {
		println("0")
	}
}

// SwitchExpression (switch 控制)
func SwitchExpression() {
	x := 100
	switch {
	case x > 0:
		println("x")
	case x < 0:
		println("-x")
	default:
		println("0")
	}
}

// ForExpressionOne (for 控制)
func ForExpressionOne() {
	for i := 0; i < 5; i++ {
		println(i)
	}
	for i := 4; i >= 0; i-- {
		println(i)
	}
}

// ForExpressionTwo (for 控制)
func ForExpressionTwo() {
	x := 0
	for x < 5 { // 相当于 while (x < 5) { ... }
		println(x)
		x++
	}
}

// ForExpressionThree (for 控制)
func ForExpressionThree() {
	x := 4
	for { // 相当于 while (true) { ... }
		println(x)
		x--

		if x < 0 {
			break
		}
	}
}

// ForExpressionFour (for 控制)
func ForExpressionFour() {
	x := []int{100, 101, 102}

	for i, n := range x {
		println(i, ":", n)
	}
}

func main() {
	IfExpression()
	SwitchExpression()
	ForExpressionOne()
	ForExpressionTwo()
	ForExpressionThree()
	ForExpressionFour()
}
