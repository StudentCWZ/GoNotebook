package main

import "fmt"

/*
	1. 在迭代期间删除或新增键值是安全的。
*/

func mapSafeOperate() {
	m := make(map[int]int)
	for i := 0; i < 10; i++ {
		m[i] = i + 10
	}
	for k := range m {
		if k == 5 {
			m[100] = 1000
		}
		delete(m, k)
		fmt.Println(k, m)
	}
}

func main() {
	mapSafeOperate()
}
