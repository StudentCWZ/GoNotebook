/*
   @Author: StudentCWZ
   @Description:
   @File: main
   @Software: GoLand
   @Project: GoNotebook
   @Date: 2022/4/6 11:11
*/

package main

/*
	1. 方法也会有同名遮蔽问题。但利用这种特性，可实现类似覆盖(override)操作。
*/

type user struct{}

type manager struct {
	user
}

func (user) toString() string {
	return "user"
}

func (m manager) toString() string {
	return m.user.toString() + "; manager"
}

func main() {
	var m manager
	println(m.toString())
	println(m.user.toString())
}
