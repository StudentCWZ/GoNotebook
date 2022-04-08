/*
   @Author: StudentCWZ
   @Description:
   @File: main
   @Software: GoLand
   @Project: GoNotebook
   @Date: 2022/4/8 09:54
*/

package main

import (
	"fmt"
)

/*
	1. 结构体(struct)可匿名嵌入其他类型。
*/

type user struct { // 结构体类型
	name string
	age  byte
}

type manager struct {
	user  // 匿名嵌入其他类型
	title string
}

func main() {
	var m manager

	m.name = "Tom" // 直接访问匿名字段的成员
	m.age = 29
	m.title = "CTO"

	fmt.Println(m)
}
