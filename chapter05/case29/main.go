package main

import (
	"fmt"
	"os"
)

/*
	1. 所谓匿名字段（anonymous field），是指没有名字，仅有类型的字段，也被称作嵌入字段或嵌入类型。
	2. 从编译器的角度看，这只是隐式地以类型名作为字段名字。可直接引用匿名字段的成员，但初始化时须当作独立字段。
*/

type attr struct {
	perm int
}

type file struct {
	name string
	attr // 仅有类型名
}

type data struct {
	os.File
}

func main() {
	f := file{
		name: "test.dat",
		attr: attr{
			perm: 0755,
		},
	}
	f.perm = 0644
	println(f.perm)

	d := data{
		File: os.File{},
	}
	fmt.Printf("%#v\n", d)
}
