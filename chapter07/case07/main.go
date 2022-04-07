/*
   @Author: StudentCWZ
   @Description:
   @File: main
   @Software: GoLand
   @Project: GoNotebook
   @Date: 2022/4/7 10:23
*/

package main

/*
	1. 接口还有一个重要的特征：将对象赋值给接口变量时，会复制该对象。
*/

type data struct {
	x int
}

func main() {
	d := data{100}
	var t interface{} = d
	println(t.(data).x)
}
