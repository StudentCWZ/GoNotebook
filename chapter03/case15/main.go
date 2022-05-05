/*
   @Author: StudentCWZ
   @Description:
   @File: main
   @Software: GoLand
   @Project: GoNotebook
   @Date: 2022/5/5 13:42
*/

package main

import "fmt"

/*
	1. 可用 for...range 完成数据迭代，支持字符串、数组指针、切片、字典、通道类型，返回索引、键值数据。
	2. 允许返回单值，或用 "_" 忽略。
	3. 无论普通 for 循环，还是 range 迭代，其定义的局部变量都会重复使用。
	4. 注意，range 会复制目标数据。受直接影响的是数组，可改用数组指针或切片类型。
	5. 如果 range 目标表达式是函数调用，也仅被执行一次。
	6. 建议嵌套循环不要超过两层，否则会难以维护。必要时可剥离，重构为函数。
*/

func data() []int {
	println("Origin data...")
	return []int{10, 20, 30}
}

func forCaseOne() {
	data := [3]string{"a", "b", "c"}
	for i, s := range data {
		println(i, s)
	}
}

func forCaseTwo() {
	data := [3]string{"a", "b", "c"}
	for i := range data {
		println(i, data[i]) // 只返回 1st value
	}

	for _, s := range data {
		println(s) // 忽略 1st value
	}

	for range data { // 仅迭代，不返回。可用来执行清空 channel 等操作
	}
}

func forCaseThree() {
	data := [3]string{"a", "b", "c"}

	for i, s := range data {
		println(&i, &s)
	}
}

func forCaseFour() {
	data := [3]int{10, 20, 30}

	for i, x := range data { // 从 data 复制品中取值
		if i == 0 {
			data[0] += 100
			data[1] += 200
			data[2] += 300
		}
		fmt.Printf("x: %d, data: %d\n", x, data[i])
	}

	for i, x := range data[:] { // 仅复制 slice，不包括底层 array
		if i == 0 {
			data[0] += 100
			data[1] += 200
			data[2] += 300
		}
		fmt.Printf("x: %d, data: %d\n", x, data[i])
	}
}

func forCaseFive() {
	for i, x := range data() {
		println(i, x)
	}
}

func main() {
	forCaseOne()
	forCaseTwo()
	forCaseThree()
	forCaseFour()
	forCaseFive()
}
