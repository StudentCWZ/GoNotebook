package main

/*
	1. 不能对 nil 字典进行写操作，但却能读
	2. 注意：内容为空的字典，与 nil 是不同的
*/

func nilMap() {
	var m map[string]int
	println(m["a"]) // 返回零值
	// m["a"] = 1 报错
}

func emptyMap() {
	var m1 map[string]int
	m2 := map[string]int{} // 已初始化，等同 make 操作
	println(m1 == nil, m2 == nil)
}

func main() {
	nilMap()
	emptyMap()
}
