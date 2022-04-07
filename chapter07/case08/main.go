/*
   @Author: StudentCWZ
   @Description:
   @File: main
   @Software: GoLand
   @Project: GoNotebook
   @Date: 2022/4/7 10:30
*/

package main

type data struct {
	x int
}

func main() {
	d := data{100}
	var t interface{} = &d
	t.(*data).x = 200
	println(t.(*data).x)
}
