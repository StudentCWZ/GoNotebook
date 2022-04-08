/*
   @Author: StudentCWZ
   @Description:
   @File: main
   @Software: GoLand
   @Project: GoNotebook
   @Date: 2022/4/8 10:04
*/

package main

import "fmt"

/*
	1. 还可以直接调用匿名字段方法，这种方式可实现与继承类似的功能。
*/

type user struct {
	name string
	age  byte
}

func (u user) ToString() string {
	return fmt.Sprintf("%+v", u)
}

type manager struct {
	user
	title string
}

func main() {
	var m manager
	m.name = "Tom"
	m.age = 29
	m.title = "CTO"
	println(m.ToString()) // 调用 user.ToString()
}
