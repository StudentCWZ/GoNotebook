package main

import "fmt"

/*
	1. 不仅仅是结构体，除接口指针和多级指针以外的任何命名类型都可以作为匿名字段
*/

type data struct {
	*int // 嵌入指针类型
	string
}

func main() {
	x := 100
	d := data{
		int:    &x,
		string: "abc", // 使用基础类型作为字段名
	}
	fmt.Printf("%#v\n", d)
}
