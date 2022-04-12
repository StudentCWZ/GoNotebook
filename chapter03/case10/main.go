/*
   @Author: StudentCWZ
   @Description:
   @File: main
   @Software: GoLand
   @Project: GoNotebook
   @Date: 2022/4/12 15:24
*/

package main

import "fmt"

/*
	1. 对于复合类型(数组、切片、字典、结构体)变量初始化时，有一些语法限制。
	   - 初始化表达式必须含类型标签
       - 左花括号必须在类型尾部，不能另起一行
       - 多个成员初始值以逗号分隔。
       - 允许多行，但每行须以逗号或右花括号结束。
*/

func main() {
	type data struct {
		x int
		s string
	}

	var a data = data{1, "abc"}
	b := data{
		1,
		"abc",
	}
	c := []int{1, 2}
	d := []int{1, 2,
		3, 4,
		5,
	}
	fmt.Println(a, b, c, d)
}
