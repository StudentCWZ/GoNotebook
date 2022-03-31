package main

/*
	1. 如果多个相同层级的匿名字段成员重名，就只能使用显式字段名访问，因为编译器无法确定目标。
	2. 严格来说，Go 并不是传统意义上的面向对象编程语言，或者说仅实现了最小面向对象机制。
	3. 匿名嵌入不是继承，无法实现多态处理。虽然配合方法集，可用接口来实现一些类似操作，但其本质是完全不同的。
*/

type file struct {
	name string
}

type log struct {
	name string
}

type data struct {
	file
	log
}

func main() {
	d := data{}
	// d.name = "name"	// 报错
	d.file.name = "file"
	d.log.name = "log"
	println(d.file.name, d.log.name)
}
