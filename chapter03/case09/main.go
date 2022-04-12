/*
   @Author: StudentCWZ
   @Description:
   @File: main
   @Software: GoLand
   @Project: GoNotebook
   @Date: 2022/4/12 10:06
*/

package main

/*
	零长度对象的地址是否相等和具体的实现版本有关，不过肯定不等于 nil。
*/

func main() {
	var a, b struct{}

	println(&a, &b)
	println(&a == &b, &a == nil)
}
