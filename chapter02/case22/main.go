/*
   @Author: StudentCWZ
   @Description:
   @File: main.go
   @Software: GoLand
   @Project: GoNotebook
   @Date: 2022/4/9 23:33
*/

package main

import "strconv"

/*
	1. 标准 strconv 可在不同进制(字符串)间转换。
*/

func main() {
	a, _ := strconv.ParseInt("1100100", 2, 32)
	b, _ := strconv.ParseInt("0144", 8, 32)
	c, _ := strconv.ParseInt("64", 16, 32)

	println(a, b, c)
	println("ob" + strconv.FormatInt(a, 2))
	println("0" + strconv.FormatInt(a, 8))
	println("Ox" + strconv.FormatInt(a, 16))
}
