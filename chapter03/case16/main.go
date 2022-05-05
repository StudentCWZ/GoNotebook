/*
   @Author: StudentCWZ
   @Description:
   @File: main
   @Software: GoLand
   @Project: GoNotebook
   @Date: 2022/5/5 14:17
*/

package main

/*
	1. 使用 goto 前，须定义标签。标签区分大小写，且未使用的标签会引发编译错误。
	2. 不能跳转到其他函数，或内层代码块内。
	3. 和 goto 定点跳转不同，break、continue 用于中断代码块执行。
	4. break: 用于 switch、 for、 select 语句，终止整个语句块执行。
	5. continue: 仅用于 for 循环、终止后续逻辑，立即进入下一轮循环
	6. 配合标签， break 和 continue 可在多层嵌套中指定目标层级。
*/

func forCaseOne() {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue // 立即进入下一轮循环
		}
		if i > 5 {
			break // 立即终止整个 for 循环
		}
		println(i)
	}
}

func forCaseTwo() {
outer:
	for x := 0; x < 5; x++ {
		for y := 0; y < 10; y++ {
			if y > 2 {
				println()
				continue outer
			}

			if x > 2 {
				break outer
			}
			print(x, ":", y, " ")
		}
	}
}

func main() {
	forCaseOne()
	forCaseTwo()
}
