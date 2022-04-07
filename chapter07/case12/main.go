/*
   @Author: StudentCWZ
   @Description:
   @File: main
   @Software: GoLand
   @Project: GoNotebook
   @Date: 2022/4/7 10:57
*/

package main

import "fmt"

/*
	1. 使用 ok-idiom 模式，即使转换失败也不会引发 panic。
	2. 还可用 switch 语句在多种类型间做出推断匹配，这样空接口就有更多发挥空间。
*/

func main() {
	var x interface{} = func(x int) string {
		return fmt.Sprintf("d: %d", x)
	}

	switch v := x.(type) {
	case nil:
		println("nil")
	case *int:
		println(*v)
	case func(int) string:
		println(v(100))
	case fmt.Stringer:
		fmt.Println(v)
	default:
		println("unknown")

	}
}
