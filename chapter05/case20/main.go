package main

import "fmt"

/*
	1. 字典（哈希表）是一种使用频率极高的数据结构。将其作为语言内置类型，从运行时层面进行优化，可获得更高效的性能。
	2. 作为无序键值对集合，字典要求 key 必须是支持相等运算符（==、!=）的数据类型，比如，数字、字符串、指针、数组、结构体，以及对应接口类型。
	3. 字典是引用类型，使用 make 函数或初始化表达式语句来创建。
	4. 访问不存在键值，默认返回零值，不会引发错误。但推荐使用 ok-idiom 模式，毕竟通过零值无法判断键值是否存在，或许存储的 value 本身就是零。
	5. 对字典的迭代，每次返回的键值次序都不相同
*/

func mapInit() {
	m := make(map[string]int)
	m["a"] = 1
	m["b"] = 2

	m2 := map[int]struct {
		x int
	}{
		1: {x: 100},
		2: {x: 200},
	}
	fmt.Println(m, m2)
}

func mapOperate() {
	m := map[string]int{
		"a": 1,
		"b": 2,
	}
	m["a"] = 10              //	修改
	m["c"] = 30              // 新增
	if v, ok := m["d"]; ok { // 使用 ok-idiom 模式判断 key 是否存在，返回值
		println(v)
	}
	delete(m, "d") // 删除键值对。不存在时，不会出错
	fmt.Println(m)
}

func mapItems() {
	m := make(map[string]int)
	for i := 0; i < 8; i++ {
		m[string('a'+i)] = i
	}
	for i := 0; i < 4; i++ {
		for k, v := range m {
			print(k, ":", v, " ")
		}
	}
	println()
}

func main() {
	mapInit()
	mapOperate()
	mapItems()
}
