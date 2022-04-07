/*
   @Author: StudentCWZ
   @Description:
   @File: main
   @Software: GoLand
   @Project: GoNotebook
   @Date: 2022/4/7 10:33
*/

package main

/*
	1. 只有当接口变量内部的两个指针(itab、data)都为 nil 时，接口才等于 nil。
*/

func main() {
	var a interface{} = nil
	var b interface{} = (*int)(nil)

	println(a == nil, b == nil)
}
