/*
   @Author: StudentCWZ
   @Description:
   @File: main
   @Software: GoLand
   @Project: GoNotebook
   @Date: 2022/4/8 09:38
*/

package main

/*
	1. 用 defer 定义定义延时调用，无论函数是否出错，它都确保结束前被调用。
*/

func test(a, b int) {
	defer println("dispose ...") // 常用来释放资源、解除锁定、或执行一些清理操作
	println(a / b)               // 可定义多个 defer， 按 FILO 顺序执行
}

func main() {
	test(10, 0)
}
