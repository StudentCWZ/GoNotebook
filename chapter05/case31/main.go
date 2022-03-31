package main

import "fmt"

/*
	1. 因未命名类型没有名字标识，自然无法作为匿名字段。
	2. 不能将基础类型和其指针类型同时嵌入，因为两者隐式名字相同。
	3. 虽然可以像普通字段那样访问匿名字段成员，但会存在重名问题。默认情况下，编译器从当前显式命名字段开始，逐步向内查找匿名字段成员。
	4. 如匿名字段成员被外层同名字段遮蔽，那么必须使用显式字段名。
*/

// type a *int
// type b **int
// type c interface{}

// type d struct {
// 	*a // 报错
// 	b  // 报错
// 	*c //报错
// }

// type data struct {
// 	*int // 不能将基础类型和其指针类型同时嵌入，因为两者隐式名字相同。
// 	int
// }

type file struct {
	name string
}

type data struct {
	file
	name string
}

func main() {
	d := data{
		name: "data",
		file: file{"file"},
	}

	d.name = "data2"
	d.file.name = "file2"
	fmt.Println(d.name, d.file.name)
}
