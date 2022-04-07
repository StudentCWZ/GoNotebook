/*
   @Author: StudentCWZ
   @Description:
   @File: main
   @Software: GoLand
   @Project: GoNotebook
   @Date: 2022/4/7 09:52
*/

package main

/*
	1. 支持匿名接口类型，可直接用于变量定义，或作为结构字段类型。
*/

type data struct{}

func (data) string() string {
	return ""
}

type node struct {
	data interface { // 匿名接口类型
		string() string
	}
}

func main() {
	var t interface { // 定义匿名接口变量
		string() string
	} = data{}
	n := node{
		data: t,
	}
	println(n.data.string())
}
