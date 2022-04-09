/*
   @Author: StudentCWZ
   @Description:
   @File: main
   @Software: GoLand
   @Project: GoNotebook
   @Date: 2022/4/9 15:09
*/

package main

import "strconv"

/*
	1. 和 Python 类似，Go 也有个名为 "_" 的特殊成员。通常作为忽略占位符使用，可作表达式左值，无法读取内容。
	2. 空标识符可用来临时规避编译器对未使用变量和导入包的错误检查。但请注意，它是预置成员，不能重新定义。
*/

func main() {
	x, _ := strconv.Atoi("12")
	println(x)
}
