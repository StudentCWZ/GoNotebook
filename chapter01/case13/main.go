/*
   @Author: StudentCWZ
   @Description:
   @File: main
   @Software: GoLand
   @Project: GoNotebook
   @Date: 2022/4/8 10:39
*/

package main

import "fmt"

/*
	1. 接口采用了 duck type 方式，也就是说无须在实现类型上添加显示声明。
	2. 另有空接口类型 interface{}，用途类似 OOP 里的 system.Object，可接收任意类型对象。
*/

type user struct {
	name string
	age  byte
}

func (u user) Print() {
	fmt.Printf("%+v\n", u)
}

type Printer interface {
	Print()
}

func main() {
	var u user

	u.name = "Tom"
	u.age = 29

	var p Printer = u
	p.Print()
}
